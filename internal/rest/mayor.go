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
// @Success      200  {object}  model.Candidate
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /mayor/current [get]
func getCurrentMayor(c *gin.Context) {
	electionPeriod, _ := mongo.GetCurrentElectionPeriod()
	if electionPeriod == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "election period not found"})
		return
	}
	c.JSON(http.StatusOK, electionPeriod.Winner)
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
	names, _ := mongo.GetAllMayorNames()
	if names == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no mayors found"})
		return
	}
	c.JSON(http.StatusOK, names)
}
