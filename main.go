package main
import (
	db "Gin/databases"
	"os"
	"github.com/gin-gonic/gin"
	"io"
)
func main()  {
	defer db.SqlDB.Close()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router :=initRouter()
	router.Run(":8000")
}
