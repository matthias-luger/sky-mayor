package rest

import (
	"net/http"
	"time"

	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/Coflnet/sky-mayor/internal/usecase"
	"github.com/gin-gonic/gin"
)

// @Summary      Get the current mayor
// @Description  Returns the current mayor
// @Tags         Mayor
// @Accept       */*
// @Produce      json
// @Success      200  {object}  model.Candidate
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /mayor/current [get]
func getCurrentMayor(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300")
	lastFetchResult := usecase.GetLastFetchResult()
	if lastFetchResult == nil {
		fetchResult, err := usecase.FetchFromHypixelApi()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.JSON(http.StatusOK, fetchResult.Mayor.Name)
		return
	}
	c.JSON(http.StatusOK, lastFetchResult.Mayor.Name)
}

// @Summary      Get names of all mayors
// @Description  Returns all mayor names
// @Tags         Mayor
// @Accept       */*
// @Produce      json
// @Success      200  {object}  []string
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /mayor/names [get]
func GetAllMayorNames(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300")
	names, _ := mongo.GetAllMayorNames()
	if names == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no mayors found"})
		return
	}
	c.JSON(http.StatusOK, names)
}

// @Summary      Get the next mayor
// @Description  Returns the mayor with the most votes in the current election. If there is currently no election, this returns null.
// @Tags         Mayor
// @Accept       */*
// @Produce      json
// @Success      200  {object}  model.Candidate
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /mayor/next [get]
func getNextMayor(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300")

	lastVoting, err := mongo.GetLastVoting()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if time.Since(lastVoting.Timestamp) > 5*time.Minute {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if lastVoting == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "last voting not found"})
		return
	}
	currentElectionPeriod, err := mongo.GetCurrentElectionPeriod()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	maxVote := lastVoting.Votes[0]
	for _, vote := range lastVoting.Votes {
		if vote.Votes > maxVote.Votes {
			maxVote = vote
		}
	}

	var nextWinner *model.Candidate
	for _, candidate := range currentElectionPeriod.Candidates {
		if candidate.Key == maxVote.MayorKey {
			nextWinner = candidate
		}
	}

	c.JSON(http.StatusOK, nextWinner)
}
