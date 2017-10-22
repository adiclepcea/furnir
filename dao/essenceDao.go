package dao

import (
	"log"
	//mysql driver
	"github.com/adiclepcea/furnir/models"
	_ "github.com/go-sql-driver/mysql"
)

//EssenceDao defines the db operations that can be done
//on an essence for persistence
type EssenceDao struct {
}

//SaveEssence will insert or update an essence
func (essenceDao EssenceDao) SaveEssence(essence models.Essence) (*models.Essence, error) {

	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if essence.ID == 0 {
		res, err := db.Exec("Insert into essences(name, code) values(?,?)", essence.Name, essence.Code)
		if err != nil {
			log.Printf("Error saving essence: %s\r\n", err.Error())
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		essence.ID = id
	} else {
		_, err = db.Exec("Update essences set name=?, code=? where essences_id=?", essence.Name, essence.Code, essence.ID)
		if err != nil {
			log.Printf("Error saving essence: %s\r\n", err.Error())
			return nil, err
		}
	}

	return &essence, nil
}

//FindEssenceByID finds the sequence with the selected id
func (essenceDao EssenceDao) FindEssenceByID(id int64) (*models.Essence, error) {
	essence := models.Essence{}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select essences_id, name, code from essences where essences_id=?", id)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&essence.ID, &essence.Name, &essence.Code)
		return &essence, nil
	}
	return nil, nil
}

//FindEssenceByName finds the sequence with the selected name
func (essenceDao EssenceDao) FindEssenceByName(name string) (*models.Essence, error) {
	essence := models.Essence{}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select essences_id, name, code from essences where name=?", name)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&essence.ID, &essence.Name, &essence.Code)
		return &essence, nil
	}
	return nil, nil
}

//FindEssenceByCode finds the sequence with the selected code
func (essenceDao EssenceDao) FindEssenceByCode(code string) (*models.Essence, error) {
	essence := models.Essence{}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select essences_id, name, code from essences where code=?", code)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&essence.ID, &essence.Name, &essence.Code)
		return &essence, nil
	}
	return nil, nil
}

//FindAllEssences returns all essences in the system
func (essenceDao EssenceDao) FindAllEssences() ([]models.Essence, error) {
	var essences []models.Essence
	essences = make([]models.Essence, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select essences_id, name, code from essences")
	if err != nil {
		return nil, err
	}
	for res.Next() {
		essence := models.Essence{}
		res.Scan(&essence.ID, &essence.Name, &essence.Code)
		essences = append(essences, essence)
	}
	return essences, nil
}

//DeleteEssenceByID deletes the essence having the passed id
func (essenceDao EssenceDao) DeleteEssenceByID(id int64) error {

	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("Delete from essences where essences_id=?", id)
	if err != nil {
		return err
	}
	return nil
}
