package models

import (
	"strconv"
	"time"
)

//ConnectionString is the string used to connect to the MySql database
var ConnectionString string

func init() {
	ConnectionString = "furnir:Furnir123@tcp(192.168.99.100:3306)/furnir?parseTime=true"
}

//Pallet is the structure used to store a Pallet in the database
type Pallet struct {
	ID      int64     `json:"id" db:"pallets_id"`
	Created time.Time `json:"created"`
}

//NewPallet will create a new Pallet
func NewPallet() (*Pallet, error) {
	db := InitDB(ConnectionString)
	defer db.Close()
	if db == nil {
		return nil, nil
	}
	res, err := db.Exec("Insert into pallets() values()")
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("Select pallets_id, created from pallets where pallets_id=" + strconv.FormatInt(id, 10))

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	pallet := Pallet{}
	err = rows.Scan(&pallet.ID, &pallet.Created)

	if err != nil {
		return nil, err
	}

	return &pallet, nil
}

//GetPalletByID obtains one pallet with a certain ID from the database
func GetPalletByID(id int64) (*Pallet, error) {
	db := InitDB(ConnectionString)
	defer db.Close()
	if db == nil {
		return nil, nil
	}

	rows, err := db.Query("Select pallets_id, created from pallets where pallets_id=" + strconv.FormatInt(id, 10))

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	pallet := Pallet{}
	err = rows.Scan(&pallet.ID, &pallet.Created)

	if err != nil {
		return nil, err
	}

	return &pallet, nil

}

//GetPallets obtains all the pallets from the database
func GetPallets() ([]Pallet, error) {
	db := InitDB(ConnectionString)
	defer db.Close()
	if db == nil {
		return nil, nil
	}

	rows, err := db.Query("Select pallets_id, created from pallets")

	if err != nil {
		return nil, err
	}

	pallets := []Pallet{}

	for rows.Next() {
		pallet := Pallet{}
		err = rows.Scan(&pallet.ID, &pallet.Created)
		if err != nil {
			return nil, err
		}
		pallets = append(pallets, pallet)
	}

	if err != nil {
		return nil, err
	}

	return pallets, nil

}
