package models

import (
	_ "encoding/json"
)

//Essence is the structure used to store an essence in the database
type Essence struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
