package databases

import (
	"database/sql"
	"log"
	_"github.com/go-sql-driver/mysql"
)
var SqlDB *sql.DB
func init()  {
	var err error
	SqlDB, err = sql.Open("mysql", "root:lichao@tcp(127.0.0.1:3306)/test?parseTime=true")

	if err != nil{
		log.Fatalln(err)
	}
	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	if err := SqlDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}
