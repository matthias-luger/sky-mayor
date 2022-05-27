package rest

import (
	_ "github.com/Coflnet/sky-mayor/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Mayor API
// @version 1.0
// @description This is the API for the Mayor votings
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
func Init() error {
	router := gin.Default()

	router.POST("/electionPeriod", postElectionPeriod)
	router.GET("/electionPeriod/:year", getElectionPeriodByYear)
	router.GET("/electionPeriod/range/:from/:to/", getElectionPeriodsByTimespan)
	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/api/doc.json")))
	return router.Run("localhost:8080")
}
