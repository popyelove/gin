package controllers

import (
	"github.com/gin-gonic/gin"
	."Gin/api/models"
	"Gin/common/tools/e"
	"Gin/common/tools/jwt"
	"net/http"
)

func GetAuthApi(c *gin.Context)  {
	username :=c.Query("username")
	password :=c.Query("password")
	auth :=Auth{Username:username,Password:password}
	data :=make(map[string]interface{})
	code :=e.INVALID_PARAMS

	isExist :=auth.CheckAuth()
	if isExist {
		token, err := jwt.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
