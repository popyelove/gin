package main
import (
	db "Gin/common/databases/mysql"
	"github.com/astaxie/beego/logs"
	routers "Gin/api/router"
)
func main()  {
	defer db.SqlDB.Close()
	logs.SetLogger(logs.AdapterFile,`{"filename":"api.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	router :=routers.InitRouter()
	router.Run(":8000")
}
