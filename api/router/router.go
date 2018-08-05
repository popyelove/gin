package router

import (
	"github.com/gin-gonic/gin"
	."Gin/api/controllers"
	"Gin/common/middleware/jwt"
)

func InitRouter() *gin.Engine  {
	router := gin.Default()
	router.GET("/auth",GetAuthApi)
	apiv1 :=router.Group("/api/v1")
	apiv1.GET("/persons",GetPersonsApi)
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/",IndexApi)
		apiv1.POST("/person",AddPersonApi)
		apiv1.GET("/person/:id",GetPersonApi)
		apiv1.PUT("/person/:id",UpdatePersonApi)
		apiv1.DELETE("/person/:id",DeletePersonApi)
	}

	return router
}
