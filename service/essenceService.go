package service

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"encoding/json"
	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/models"
)


//Error structure used to pass errors
type Error struct {
	Message string `json:"message"`
}

//EssenceService serves as a prototype for serving essence
//operations
type EssenceService struct{
	essenceRepo dao.EssenceDao
}

//NewEssenceService returns a new instance of EssenceService using the passed
//in dao.EssenceDao
func NewEssenceService(essenceRepo dao.EssenceDao)(EssenceService){
	es := EssenceService{essenceRepo: essenceRepo}
	return es
}

//PutEssence translates a PUT request to an essence modification
func (essenceService *EssenceService) PutEssence(w http.ResponseWriter, r *http.Request){
	var essenceRepo dao.EssenceDao

	decoder := json.NewDecoder(r.Body)
	essence := models.Essence{}

	encoder := json.NewEncoder(w)
	
	err := decoder.Decode(&essence)

	if err!=nil || essence.ID == 0{
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid essence supplied"})
		return
	}
	
	ess, err := essenceRepo.SaveEssence(essence)
	
	encoder.Encode(ess)
}

//DeleteEssence translates a DELETE request to an essence deletion
func (essenceService *EssenceService) DeleteEssence(w http.ResponseWriter, r *http.Request){
	var essenceRepo dao.EssenceDao

	encoder := json.NewEncoder(w)

	strID := r.URL.Query().Get("id")

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID,10,64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting an essence \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		err = essenceRepo.DeleteEssenceByID(id)

		if err != nil {
			log.Printf("Error deleting essence: %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Error deleting essence: %s", err.Error())})
			return
		}
	}

}

//PostEssence translates a POST request to an essence creation
func (essenceService *EssenceService) PostEssence(w http.ResponseWriter, r *http.Request){
	var essenceRepo dao.EssenceDao

	decoder := json.NewDecoder(r.Body)
	essence := models.Essence{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&essence)

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid essence supplied"})
		return
	}

	essence.ID = 0
	
	ess, err := essenceRepo.SaveEssence(essence)

	encoder.Encode(ess)
}

//GetEssence translates a GET request to an essence request from the db
func (essenceService *EssenceService) GetEssence(w http.ResponseWriter, r *http.Request) {
	
	var essenceRepo dao.EssenceDao

	strID := r.URL.Query().Get("id")
	code := r.URL.Query().Get("code")
	name := r.URL.Query().Get("name")

	encoder := json.NewEncoder(w)

	var essence *models.Essence
	var err error

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID,10,64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting an essence \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		essence, err = essenceRepo.FindEssenceByID(id)

		if err != nil {
			log.Printf("Error retrieving essence with id %d: %s \r\n", id, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})			
			return
		}

	} else if len(code) != 0 {
		essence, err = essenceRepo.FindEssenceByCode(code)

		if err != nil {
			log.Printf("Error retrieving essence with code %s: %s \r\n", code, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}
	} else if len(name) != 0 {
		essence, err = essenceRepo.FindEssenceByName(name)

		if err != nil {
			log.Printf("Error retrieving essence with name %s: %s \r\n", name, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}
	} else {
		//return all essences
		essences, err := essenceRepo.FindAllEssences()

		if err != nil {
			log.Printf("Error retrieving essences %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

		encoder.Encode(essences)
		return
	}

	if essence != nil {
		encoder.Encode(*essence)
	}

}