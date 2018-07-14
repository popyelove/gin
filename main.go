package main
import (
	db "Gin/databases"
)
func main()  {
	defer db.SqlDB.Close()
	router :=initRouter()
	router.Run(":8000")
}
