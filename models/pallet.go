package models

import (
	_"encoding/json"
)

//Pallet is the structure used to store a Pallet in the database
type Pallet struct{
	ID int64 `json:"id"`
	Code string `json:"code"`
}

//NewPallet will create a new Pallet
func NewPallet()(*Pallet, error){	
	db := InitDB()
	defer db.Close()
	if db==nil{
		return nil, nil
	}
	res, err := db.Exec("Insert into pallets() values()")
	if err!=nil {
		return nil,err
	}
	id, err := res.LastInsertId()
	if err!=nil {
		return nil,err
	}

	return &Pallet{ID:id,Code:""},nil
}
