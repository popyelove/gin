package mysql

import (
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
	"Gin/common/tools/setting"
	"fmt"
)
var SqlDB *sql.DB
func init()  {
	var err error
	connstr :=setting.DatabaseSetting.User+":"+setting.DatabaseSetting.Password+"@tcp("+setting.DatabaseSetting.Host+")"+"/test?parseTime=true"
	fmt.Print(connstr)
	SqlDB, err = sql.Open("mysql",connstr)

	if err != nil{
		log.Fatalln(err)
	}
	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	if err := SqlDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}
