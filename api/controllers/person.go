package controllers

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	."Gin/api/models"
	"fmt"
)

func IndexApi(c *gin.Context)  {
	c.String(http.StatusOK, "It works")
}
func AddPersonApi(c *gin.Context)  {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person :=Person{FirstName:firstName,LastName:lastName}
	ra_rows, err := person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra_rows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func GetPersonsApi(c *gin.Context)  {
	person :=Person{}
	persons,err :=person.GetPersons()
	if err!=nil{
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
func GetPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err!=nil {
		log.Fatalln(err)
	}
	person :=Person{Id:id}
	persons,err :=person.GetPerson()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": persons,
	})
}
func UpdatePersonApi(c *gin.Context)  {
	cid := c.Param("id")
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(cid)
	if err!=nil {
		log.Fatalln(err)
	}
	person := Person{Id: id,FirstName:firstName,LastName:lastName}
	nums,err:=person.UpdatePerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update person %d successful %d",id,nums)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func DeletePersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	person := Person{Id: id}
	ra, err := person.DeletePerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
