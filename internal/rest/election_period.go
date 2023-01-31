package rest

import (
	"net/http"
	"strconv"

	"github.com/Coflnet/sky-mayor/internal/model"
	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @Summary      Get election periods by timespan
// @Description  Returns all election periods that took place in a given timespan
// @Tags         Election periods
// @Accept       */*
// @Produce      json
// @Param        from  query     int  true  "from The beginning of the selected timespan"  Format(int64)
// @Param        to    query     int  true  "The end of the selected timespan"             Format(int64)
// @Success      200   {object}  []model.ElectionPeriod
// @Failure      400   {object}  nil
// @Failure      404   {object}  nil
// @Router       /electionPeriod/range [get]
func getElectionPeriodsByTimespan(c *gin.Context) {
	fromParam := c.Query("from")
	toParam := c.Query("to")

	from, err := strconv.ParseInt(fromParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expection param 'from' to be type int"})
		return
	}

	to, err := strconv.ParseInt(toParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expection param 'to' to be type int"})
		return
	}

	electionPeriod, _ := mongo.GetElectionPeriodsByTimespan(from/1000, to/1000)
	if electionPeriod == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "election period not found"})
		return
	}
	c.JSON(http.StatusOK, electionPeriod)
}

// @Summary      Get the election period of a certain year
// @Description  Returns the election periods that took place in a given year
// @Tags         Election periods
// @Accept       */*
// @Produce      json
// @Param        year  path      int  true  "the searched year"
// @Success      200   {object}  model.ElectionPeriod
// @Failure      400   {object}  nil
// @Failure      404   {object}  nil
// @Router       /electionPeriod/{year} [get]
func getElectionPeriodByYear(c *gin.Context) {
	yearParam := c.Param("year")

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "expection param 'year' to be type int"})
		return
	}

	electionPeriod, _ := mongo.GetElectionPeriodByYear(year)
	if electionPeriod == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "election period not found"})
		return
	}
	c.JSON(http.StatusOK, electionPeriod)

}

// @Summary      Inserts election periods
// @Description  Endpoint to insert election periods, should only be used to insert missing/hisotical data
// @Tags         Election periods
// @Accept       */*
// @Produce      json
// @Param        periods  body      []model.ElectionPeriod  true  "the election periods that are going to be inserted"
// @Success      201      {object}  []model.ElectionPeriod
// @Failure      400      {object}  nil
// @Router       /electionPeriod [post]
func postElectionPeriod(c *gin.Context) {

	var newElectionPeriods = []*model.ElectionPeriod{}

	if err := c.BindJSON(&newElectionPeriods); err != nil {
		log.Error().Err(err).Msg("error parsing json")
		c.JSON(http.StatusBadRequest, gin.H{"message": "error parsing json"})
		return
	}

	if err := mongo.InsertElectionPeriods(newElectionPeriods); err != nil {
		log.Error().Err(err).Msg("error inserting election period")
		c.JSON(http.StatusBadRequest, gin.H{"message": "error inserting election period"})
		return
	}
	c.JSON(http.StatusCreated, newElectionPeriods)
}

