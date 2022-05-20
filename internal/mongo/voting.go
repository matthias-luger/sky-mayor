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

func InsertVoting(voting *model.Voting) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	voting.ID = primitive.NewObjectID()

	result, err := votingCollection.InsertOne(ctx, voting)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error when updating voting with id %s", voting.ID)
		return err
	}

	log.Info().Msgf("%v id, voting (timestamp %s) was inserted", result.InsertedID, voting.Timestamp.Format("2006-01-02 15:04:05"))
	return nil
}

func GetLastVoting() (*model.Voting, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cur, err := votingCollection.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{"timestamp", -1}}).SetLimit(1))

	if err != nil {
		log.Error().Err(err).Msgf("error finding voting data")
		return nil, err
	}

	var result = &model.Voting{}

	if cur.Next(ctx) {

		err := cur.Decode(result)
		if err != nil {
			log.Error().Err(err).Msgf("error decoding voting data")
			return nil, err
		}

		log.Info().Msgf("successfully found last voting data (timestamp %s)", result.Timestamp.Format("2006-01-02 15:04:05"))

	} else {
		result = nil
		log.Info().Msgf("no voting data found")
	}

	return result, nil
}
