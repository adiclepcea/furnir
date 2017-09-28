package models

import (
	"database/sql"
	"log"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//InitDB initializes the connection to the database
//example ConnectionString "furnir:Furnir123@tcp(127.0.0.1:13306)/furnir?parseTime=true"
func InitDB(ConnectionString string) *sql.DB {
	//"user:pass@/db"
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("Connected to DB")
	return db
}
