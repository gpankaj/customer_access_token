package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gpankaj/storage_access_tokken/src/http"
	"github.com/gpankaj/storage_access_tokken/src/repository/db"
	"github.com/gpankaj/storage_access_tokken/src/repository/rest"
	"github.com/gpankaj/storage_access_tokken/src/services/access_token_service"
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