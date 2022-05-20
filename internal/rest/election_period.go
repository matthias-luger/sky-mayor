package rest

import (
	"net/http"
	"strconv"

	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InitElectionPeriodREST() {
	router := gin.Default()
	router.POST("/electionPeriod", postElectionPeriod)
	router.GET("/electionPeriod/:year", getElectionPeriodByYear)
	router.Run("localhost:8080")
}

func getElectionPeriodByYear(c *gin.Context) {
	yearParam := c.Param("year")

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "expection param year to be type int"})
		return
	}

	electionPeriod, _ := mongo.GetElectionPeriodByYear(year)
	if electionPeriod == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "election period not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, electionPeriod)

}

func postElectionPeriod(c *gin.Context) {

	var newElectionPeriods = []*model.ElectionPeriod{}

	if err := c.BindJSON(&newElectionPeriods); err != nil {
		log.Error().Err(err).Msg("error parsing json")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error parsing json"})
		return
	}

	if err := mongo.InsertElectionPeriods(newElectionPeriods); err != nil {
		log.Error().Err(err).Msg("error inserting election period")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error inserting election period"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newElectionPeriods)
}
