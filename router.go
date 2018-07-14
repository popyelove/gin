package main

import (
	"github.com/gin-gonic/gin"
	."Gin/apis"
)

func initRouter() *gin.Engine  {
	router := gin.Default()
	router.GET("/",IndexApi)
	router.POST("/person",AddPersonApi)
	router.GET("/persons",GetPersonsApi)
	router.GET("/person/:id",GetPersonApi)
	router.PUT("/person/:id",UpdatePersonApi)
	router.DELETE("/person/:id",DeletePersonApi)
	return router
}
