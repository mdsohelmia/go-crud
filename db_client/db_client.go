package dbclient

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)
var DbClient *sql.DB

func InitilizeDbConnection() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kodezam_v2")

	if err != nil {
		panic(err.Error())
	}
	pingError :=db.Ping()

	if pingError !=nil{
		panic(pingError.Error())
	}
	DbClient =db;
}