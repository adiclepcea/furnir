package dao

import (
	"database/sql"
	"log"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//InitDB initializes the connection to the database
func InitDB() (*sql.DB, error) {
	//"user:pass@/db"
	db, err := sql.Open("mysql", "furnir:Furnir123@tcp(127.0.0.1:3306)/furnir")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}
