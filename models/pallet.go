package models

import (
	_ "encoding/json"
	"time"
)

//ConnectionString is the string used to connect to the MySql database
var ConnectionString string

func init() {
	ConnectionString = "furnir:Furnir123@tcp(192.168.99.100:3306)/furnir?parseTime=true"
}

//Pallet is the structure used to store a Pallet in the database
type Pallet struct {
	ID      int64     `json:"id"`
	Created time.Time `json:"created"`
	Essence Essence   `json:"essence"`
}

