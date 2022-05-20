package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Coflnet/sky-mayor/internal/metrics"
	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/rs/zerolog/log"
)

var currYear int

func StartFetch() {
	if err := Fetch(); err != nil {
		log.Error().Err(err).Msg("error while fetching")
		metrics.AddError()
	}
	for range time.Tick(time.Minute * 5) {
		if err := Fetch(); err != nil {
			log.Error().Err(err).Msg("error while fetching")
			metrics.AddError()
		}
	}
}

func Fetch() error {
	apiResponse, err := FetchFromHypixelApi()

	if err != nil {
		return err
	}

	if isApiDataAlreadyPresent(apiResponse) {
		log.Info().Msg("api data is already present, retry in 5min")
		return nil
	}

	var votes = []*model.Vote{}
	for i := 0; i < len(apiResponse.Current.Candidates); i++ {
		votes = append(votes, &model.Vote{
			MayorKey: apiResponse.Current.Candidates[i].Key,
			Votes:    apiResponse.Current.Candidates[i].Votes,
		})
	}

	err = mongo.InsertVoting(&model.Voting{
		Year:      apiResponse.Current.Year,
		Votes:     votes,
		Timestamp: time.Unix(apiResponse.LastUpdated/1000, 0),
	})

	if err != nil {
		log.Error().Err(err).Msg("Error inserting voting data")
		return err
	}
	metrics.VotingInserted()

	if currYear == 0 || currYear != apiResponse.Current.Year {
		if err = createCurrentElectionPeriodIfNeeded(apiResponse); err != nil {
			log.Error().Err(err).Msg("Error checking/creating curring election period")
			return err
		}
		if err = updateWinnerOfLastYear(apiResponse); err != nil {
			log.Error().Err(err).Msgf("Error updating winner of last year (%d)", apiResponse.Mayor.Election.Year)
			return err
		}
		currYear = apiResponse.Current.Year
	}
	return nil
}

func createCurrentElectionPeriodIfNeeded(apiResponse *model.ApiElectionResponse) error {
	election, err := mongo.GetElectionPeriodByYear(apiResponse.Current.Year)
	if err != nil {
		return err
	}
	if election != nil {
		return nil
	}

	start, end := GetTimeSpanForSkyblockYear(apiResponse.Current.Year)
	err = mongo.InsertElectionPeriod(&model.ElectionPeriod{
		Year:       apiResponse.Current.Year,
		Start:      start,
		End:        end,
		Candidates: getCandidatesFromApiCandidates(apiResponse.Current.Candidates),
		Winner:     nil,
	})

	if err != nil {
		log.Error().Err(err).Msgf("Error inserting current election period (year %d)", apiResponse.Current.Year)
		return err
	}
	metrics.ElectionPeriodInserted()
	return nil
}

func updateWinnerOfLastYear(apiResponse *model.ApiElectionResponse) error {
	lastElectionPeriod, _ := mongo.GetElectionPeriodByYear(apiResponse.Mayor.Election.Year)
	winner := &model.Candidate{
		Name:  apiResponse.Mayor.Name,
		Key:   apiResponse.Mayor.Key,
		Perks: apiResponse.Mayor.Perks,
	}

	if lastElectionPeriod == nil {
		start, end := GetTimeSpanForSkyblockYear(apiResponse.Mayor.Election.Year)
		lastElectionPeriod = &model.ElectionPeriod{
			Year:       apiResponse.Mayor.Election.Year,
			Start:      start,
			End:        end,
			Candidates: getCandidatesFromApiCandidates(apiResponse.Mayor.Election.Candidates),
			Winner:     winner,
		}
		err := mongo.InsertElectionPeriod(lastElectionPeriod)

		if err != nil {
			log.Error().Err(err).Msgf("Error inserting last election period (year %d)", apiResponse.Mayor.Election.Year)
			return err
		}
		metrics.ElectionPeriodInserted()

		return nil
	}

	lastElectionPeriod.Winner = winner
	mongo.UpdateElectionPeriod(lastElectionPeriod)
	return nil
}

func getCandidatesFromApiCandidates(apiCandidates []*model.ApiCandidates) []*model.Candidate {
	var candidates = []*model.Candidate{}
	for i := 0; i < len(apiCandidates); i++ {
		candidate := model.Candidate{
			Name:  apiCandidates[i].Name,
			Key:   apiCandidates[i].Key,
			Perks: apiCandidates[i].Perks,
		}
		candidates = append(candidates, &candidate)
	}
	return candidates
}

func isApiDataAlreadyPresent(apiResponseData *model.ApiElectionResponse) bool {
	lastVoting, _ := mongo.GetLastVoting()
	if lastVoting == nil {
		return false
	}
	return lastVoting.Timestamp.Unix() == apiResponseData.LastUpdated/1000
}

func FetchFromHypixelApi() (*model.ApiElectionResponse, error) {
	url := "https://api.hypixel.net/resources/skyblock/election"

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating request")
		return nil, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Error().Err(getErr).Msgf("Error fetching data from %s", url)
		return nil, err
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Error().Err(readErr).Msg("Error reading response body")
		return nil, err
	}

	apiResponse := &model.ApiElectionResponse{}

	jsonErr := json.Unmarshal(body, apiResponse)
	if jsonErr != nil {
		log.Error().Err(jsonErr).Msg("Error parsing response")
		return nil, err
	}

	return apiResponse, nil
}
