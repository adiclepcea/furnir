package dao

import (
	"database/sql"
	"log"
	//mysql driver
	"github.com/adiclepcea/furnir/models"
)

//PalletDao defines the db operations that can be done
//on a pallet for persistence
type PalletDao struct {
}

//SavePallet will insert or update a pallet
func (palletDao PalletDao) SavePallet(pallet models.Pallet) (*models.Pallet, error) {
	log.Printf("Save pallet: %d\r\n", pallet.ID)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if pallet.ID == 0 {
		var res sql.Result
		if pallet.Essence.ID != 0 {
			res, err = db.Exec("Insert into pallets(essences_id) values(?)", pallet.Essence.ID)
		} else {
			res, err = db.Exec("Insert into pallets() values()")
		}
		if err != nil {
			log.Printf("Error saving pallet: %s\r\n", err.Error())
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		pallet.ID = id

	} else {
		_, err = db.Exec("Update pallets set essences_id=? where pallets_id=?", pallet.Essence.ID, pallet.ID)
		if err != nil {
			log.Printf("Error saving pallet: %s\r\n", err.Error())
			return nil, err
		}
	}

	return palletDao.FindPalletByID(pallet.ID)
}

//FindPalletByID finds the pallet with the selected id
func (palletDao PalletDao) FindPalletByID(id int64) (*models.Pallet, error) {
	log.Printf("Find pallet by id: %d\r\n", id)
	pallet := models.Pallet{}
	pallet.Essence = models.Essence{}
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select p.pallets_id, p.created, e.essences_id,e.name, e.code  from pallets p left outer join essences e on p.essences_id=e.essences_id where pallets_id=?", id)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&pallet.ID, &pallet.Created, &pallet.Essence.ID, &pallet.Essence.Name, &pallet.Essence.Code)
		return &pallet, nil
	}
	return nil, nil
}

//FindPalletsByEssenceID finds the pallets with the selected id
func (palletDao PalletDao) FindPalletsByEssenceID(id int64) ([]models.Pallet, error) {
	log.Printf("Find pallet by essence id: %d\r\n", id)
	var pallets []models.Pallet
	pallets = make([]models.Pallet, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select p.pallets_id, p.created, e.essences_id,e.name, e.code  from pallets p inner join essences e on p.essences_id=e.essences_id where e.essences.id=?", id)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		pallet := models.Pallet{}
		pallet.Essence = models.Essence{}
		res.Scan(&pallet.ID, &pallet.Created, &pallet.Essence.ID, &pallet.Essence.Name, &pallet.Essence.Code)
		pallets = append(pallets, pallet)
	}
	return pallets, nil
}

//FindPalletsByEssenceCode finds the pallets with the selected essence code
func (palletDao PalletDao) FindPalletsByEssenceCode(code string) ([]models.Pallet, error) {
	log.Printf("Find pallet by essence code: %s\r\n", code)
	var pallets []models.Pallet
	pallets = make([]models.Pallet, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select p.pallets_id, p.created, e.essences_id,e.name, e.code  from pallets p inner join essences e on p.essences_id=e.essences_id where e.essences.code=?", code)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		pallet := models.Pallet{}
		pallet.Essence = models.Essence{}
		res.Scan(&pallet.ID, &pallet.Created, &pallet.Essence.ID, &pallet.Essence.Name, &pallet.Essence.Code)
		pallets = append(pallets, pallet)
	}
	return pallets, nil
}

//FindPalletsByEssenceName finds the pallets with the selected essence name
func (palletDao PalletDao) FindPalletsByEssenceName(name string) ([]models.Pallet, error) {
	log.Printf("Find pallets by essence name: %s\r\n", name)
	var pallets []models.Pallet
	pallets = make([]models.Pallet, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select p.pallets_id, p.created, e.essences_id,e.name, e.code  from pallets p inner join essences e on p.essences_id=e.essences_id where e.essences.name=?", name)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		pallet := models.Pallet{}
		pallet.Essence = models.Essence{}
		res.Scan(&pallet.ID, &pallet.Created, &pallet.Essence.ID, &pallet.Essence.Name, &pallet.Essence.Code)
		pallets = append(pallets, pallet)
	}
	return pallets, nil
}

//FindAllPallets returns all essences in the system
func (palletDao PalletDao) FindAllPallets() ([]models.Pallet, error) {
	log.Println("Find all pallets")
	var pallets []models.Pallet
	pallets = make([]models.Pallet, 0)
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	res, err := db.Query("Select p.pallets_id, p.created, e.essences_id,e.name, e.code  from pallets p left outer join essences e on p.essences_id=e.essences_id")
	if err != nil {
		return nil, err
	}
	for res.Next() {
		pallet := models.Pallet{}
		pallet.Essence = models.Essence{}
		res.Scan(&pallet.ID, &pallet.Created, &pallet.Essence.ID, &pallet.Essence.Name, &pallet.Essence.Code)
		pallets = append(pallets, pallet)
	}
	return pallets, nil
}

//DeletePalletByID deletes the essence having the passed id
func (palletDao PalletDao) DeletePalletByID(id int64) error {
	log.Printf("Delete pallet by id: %d\r\n", id)
	db, err := InitDB()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("Delete from pallets where pallets_id=?", id)
	if err != nil {
		return err
	}
	return nil
}
