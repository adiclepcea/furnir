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
	ID   int64  `json:"id"`
	Created time.Time `json:"created"`
	Essence Essence `json:"essence"`
}

//NewPallet will create a new Pallet
/*func NewPallet() (*Pallet, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
>>>>>>> Add db structure and essence REST
	}
	defer db.Close()
	res, err := db.Exec("Insert into pallets() values()")
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err

	}

	return &Pallet{ID: id, Code: ""}, nil
}*/
/*
//MovePieceToPallet will move the specified piece from
//this container the specified one
func MovePieceToPallet(piece ScannedPiece, toPallet Pallet) error {
	db, err := InitDB()
	if err != nil {
		return nil
	}

	defer db.Close()

	res, err := db.exec("select ")
}*/
