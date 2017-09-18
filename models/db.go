package models

import 
(
	"database/sql"
	"log"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//InitDB initializes the connection to the database
func InitDB(){
	//"user:pass@/db"
	db, err := sql.Open("mysql", "furnir:Furnir123@tcp(127.0.0.1:13306)/furnir")
	if err!=nil {
		log.Panic(err)
	}

	if err=db.Ping(); err!=nil{
		log.Panic(err)
	}

	log.Println("Connected to DB")
}