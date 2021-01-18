package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/customer_access_token/src/http"
	"github.com/gpankaj/customer_access_token/src/repository/db"

	"github.com/gpankaj/customer_access_token/src/repository/rest"
	"github.com/gpankaj/customer_access_token/src/services/access_token_service"
)

var (
	router = gin.Default()
)




func StartApplication() {

	//Get repository you will use
	dbRepository := db.NewRepository()

	atService:= access_token_service.NewService(rest.NewRepository(), dbRepository)

	atHandler := http.NewHandler(atService)
	router.Use(cors.Default())
	router.GET("/oauth/access_token/:access_token_id",atHandler.GetById)

	router.POST("/oauth/access_token/",atHandler.Create)
	router.Run(":9091")
}