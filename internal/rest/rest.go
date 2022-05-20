package rest

import "github.com/gin-gonic/gin"

func Init() error {
	router := gin.Default()
	router.POST("/electionPeriod", postElectionPeriod)
	router.GET("/electionPeriod/:year", getElectionPeriodByYear)
	router.GET("/electionPeriod/range/:from/:to/", getElectionPeriodsByTimespan)
	return router.Run("localhost:8080")
}
