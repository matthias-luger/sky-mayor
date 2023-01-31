package rest

import (
	"net/http"

	"github.com/Coflnet/sky-mayor/internal/mongo"
	"github.com/gin-gonic/gin"
)

// @Summary      Get the current mayor
// @Description  Returns the current mayor
// @Tags         Mayor
// @Accept       */*
// @Produce      json
// @Success      200   {object}  model.Candidate
// @Failure      400   {object}  nil
// @Failure      404   {object}  nil
// @Router       /mayor/current [get]
func getCurrentMayor(c *gin.Context) {
	electionPeriod, _ := mongo.GetCurrentElectionPeriod()
	if electionPeriod == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "election period not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, electionPeriod.Winner)
}
