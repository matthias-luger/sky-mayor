package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/rs/zerolog/log"
)

var currYear int

func StartFetch() {
	Fetch()
	for range time.Tick(time.Minute * 5) {
		Fetch()
	}
}

func Fetch() {
	apiResponse := FetchFromHypixelApi()

	if isApiDataAlreadyPresent(apiResponse) {
		log.Info().Msg("api data is already present, retry in 5min")
		return
	}

	var votes = []*model.Vote{}
	for i := 0; i < len(apiResponse.Current.Candidates); i++ {
		votes = append(votes, &model.Vote{
			MayorKey: apiResponse.Current.Candidates[i].Key,
			Votes:    apiResponse.Current.Candidates[i].Votes,
		})
	}

	mongo.InsertVoting(&model.Voting{
		Year:      apiResponse.Current.Year,
		Votes:     votes,
		Timestamp: time.Unix(apiResponse.LastUpdated/1000, 0),
	})

	if currYear == 0 || currYear != apiResponse.Current.Year {
		createCurrentElectionPeriodIfNeeded(apiResponse)
		updateWinnerOfLastYear(apiResponse)
		currYear = apiResponse.Current.Year
	}
}

func createCurrentElectionPeriodIfNeeded(apiResponse model.ApiElectionResponse) {
	election, _ := mongo.GetElectionPeriodByYear(apiResponse.Current.Year)
	if election != nil {
		return
	}

	start, end := GetTimeSpanForSkyblockYear(apiResponse.Current.Year)
	mongo.InsertElectionPeriod(&model.ElectionPeriod{
		Year:       apiResponse.Current.Year,
		Start:      start,
		End:        end,
		Candidates: getCandidatesFromApiCandidates(apiResponse.Current.Candidates),
		Winner:     nil,
	})
}

func updateWinnerOfLastYear(apiResponse model.ApiElectionResponse) {
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
		mongo.InsertElectionPeriod(lastElectionPeriod)
		return
	}

	lastElectionPeriod.Winner = winner
	mongo.UpdateElectionPeriod(lastElectionPeriod)
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

func isApiDataAlreadyPresent(apiResponseData model.ApiElectionResponse) bool {
	lastVoting, _ := mongo.GetLastVoting()
	if lastVoting == nil {
		return false
	}
	return lastVoting.Timestamp.Unix() == apiResponseData.LastUpdated/1000
}

func FetchFromHypixelApi() model.ApiElectionResponse {
	url := "https://api.hypixel.net/resources/skyblock/election"

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating request")
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal().Err(getErr).Msgf("Error fetching data from %s", url)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal().Err(readErr).Msg("Error reading response body")
	}

	apiResponse := model.ApiElectionResponse{}

	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		log.Fatal().Err(jsonErr).Msg("Error parsing response")
	}

	return apiResponse
}
