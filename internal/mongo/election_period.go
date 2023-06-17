package mongo

import (
	"context"
	"time"

	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertElectionPeriod(electionPeriod *model.ElectionPeriod) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	electionPeriod.ID = primitive.NewObjectID()

	result, err := electionPeriodCollection.InsertOne(ctx, electionPeriod)
	if err != nil {
		log.Error().Err(err).Msgf("%v id, there was an error when trying to insert election period with id %d", electionPeriod.ID, electionPeriod.Year)
		return err
	}

	log.Info().Msgf("%v id, election period %d inserted", result.InsertedID, electionPeriod.Year)
	return nil
}

func InsertElectionPeriods(electionPeriods []*model.ElectionPeriod) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var insertList []interface{}
	for i := 0; i < len(electionPeriods); i++ {
		electionPeriods[i].ID = primitive.NewObjectID()
		insertList = append(insertList, electionPeriods[i])
	}

	result, err := electionPeriodCollection.InsertMany(ctx, insertList)
	if err != nil {
		log.Error().Err(err).Msgf("There was an error when trying to insert multiple election periods. Inserts: %v", electionPeriods)
		return err
	}

	log.Info().Msgf("inserted multple election periods: ", result.InsertedIDs)
	return nil
}

func UpdateElectionPeriod(electionPeriod *model.ElectionPeriod) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, err := electionPeriodCollection.ReplaceOne(ctx, bson.D{{"_id", electionPeriod.ID}}, electionPeriod)

	if err != nil {
		log.Error().Err(err).Msgf("there was an error when updating eleciton period with id %s", electionPeriod.ID)
		return err
	}

	log.Info().Msgf("id: %v, successfully replaced election period of year %d", electionPeriod.ID, electionPeriod.Year)
	return nil
}

func GetElectionPeriodByYear(year int) (*model.ElectionPeriod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cur, err := electionPeriodCollection.Find(ctx, bson.D{{"year", year}}, options.Find().SetLimit(1))

	if err != nil {
		log.Error().Err(err).Msgf("error finding election period for year %d", year)
		return nil, err
	}

	var result = &model.ElectionPeriod{}

	if cur.Next(ctx) {

		log.Info().Msgf("successfully found election period for year %d", year)

		err := cur.Decode(result)
		if err != nil {
			log.Error().Err(err).Msgf("error decoding election period for year %d", year)
			return nil, err
		}
	} else {
		result = nil
		log.Info().Msgf("no election period found for year %d", year)
	}

	return result, nil
}

func GetElectionPeriodsByTimespan(from int64, to int64) ([]*model.ElectionPeriod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cur, err := electionPeriodCollection.Find(ctx, bson.M{
		"$and": []bson.M{
			{"start": bson.M{"$gte": time.Unix(from, 0)}},
			{"end": bson.M{"$lte": time.Unix(to, 0)}},
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error finding election period from %d to %d", from, to)
		return nil, err
	}

	var resultList = []*model.ElectionPeriod{}

	for cur.Next(ctx) {

		var result = &model.ElectionPeriod{}
		err := cur.Decode(result)
		if err != nil {
			log.Error().Err(err).Msgf("error decoding election period")
			return nil, err
		}
		resultList = append(resultList, result)
	}

	log.Info().Msgf("successfully found %d election periods from %d to %d", len(resultList), from, to)

	return resultList, nil
}

func GetCurrentElectionPeriod() (*model.ElectionPeriod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var result = &model.ElectionPeriod{}

	opts := options.FindOne().SetSort(bson.D{{"year", -1}})
	err := electionPeriodCollection.FindOne(ctx, bson.D{}, opts).Decode(result)

	if err != nil {
		log.Error().Err(err).Msgf("error finding current election period")
		return nil, err
	}

	log.Info().Msgf("successfully found current election period %d", result.Year)

	return result, nil
}

func GetPreviouslyElectedMayor() (*model.Candidate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var result = &model.ElectionPeriod{}

	opts := options.FindOne().SetSort(bson.M{"end": -1})
	err := electionPeriodCollection.FindOne(ctx, bson.M{"end": bson.M{"$lt": time.Now()}}, opts).Decode(result)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("%d", result.Year)

	return result.Winner, nil
}

func GetAllMayorNames() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cursor, err := electionPeriodCollection.Aggregate(ctx, []bson.M{
		{
			"$unwind": "$candidates",
		},
		{
			"$group": bson.M{
				"_id": "$candidates.name",
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error finding mayors")
		return nil, err
	}

	var results []string
	var result bson.M
	for cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Error().Err(err).Msgf("error decoding mayor name")
			continue
		}
		if result["_id"] != nil {
			results = append(results, result["_id"].(string))
		}
	}

	log.Info().Msg("successfully found all distinct mayor names")

	return results, nil
}
