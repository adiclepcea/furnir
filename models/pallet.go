package models

import (
	_"encoding/json"
)

//Pallet is the structure used to store a Pallet in the database
type Pallet struct{
	ID int `json:"id"`
	Code string `json:"code"`
}

//New will create a new Pallet
func (palet Pallet) New()(Pallet){	
	return Pallet{ID:0,Code:""}
}
