package dao

import (
	"database/sql"
	"fmt"
	"log"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Database means the name of the database to be used
var Database = "furnir"

//DbUser is the user to use when connecting to the database
var DbUser = "furnir"

//DbPassword is the password to use when connecting to the database
var DbPassword = "Furnir123"

//DbServer is the server to be used when connecting to the database
var DbServer = "localhost"

//DbPort is an integer that designates the port to use when connecting to the database
var DbPort = 3306

//InitDB initializes the connection to the database
func InitDB() (*sql.DB, error) {
	//"user:pass@/db"
	//db, err := sql.Open("mysql", "furnir:Furnir123@tcp(127.0.0.1:3306)/furnir?parseTime=true")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", DbUser, DbPassword, DbServer, DbPort, Database))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
