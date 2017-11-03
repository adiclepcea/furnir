package service

import (
	"encoding/json"
	"fmt"
	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/models"
	"github.com/adiclepcea/furnir/pdf"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

//PalletService serves as a prototype for serving pallet
//operations
type PalletService struct {
	palletRepo dao.PalletDao
}

//NewPalletService returns a new instance of PalletService using the passed
//in dao.PalletDao
func NewPalletService(palletRepo dao.PalletDao) PalletService {
	es := PalletService{palletRepo: palletRepo}
	return es
}

//PutPallet translates a PUT request to an pallet modification
func (palletService *PalletService) PutPallet(w http.ResponseWriter, r *http.Request) {
	var palletRepo dao.PalletDao

	decoder := json.NewDecoder(r.Body)
	pallet := models.Pallet{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&pallet)

	if err != nil || pallet.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid pallet supplied"})
		return
	}

	pal, err := palletRepo.SavePallet(pallet)

	encoder.Encode(pal)
}

//DeletePallet translates a DELETE request to an pallet deletion
func (palletService *PalletService) DeletePallet(w http.ResponseWriter, r *http.Request) {
	var palletRepo dao.PalletDao

	encoder := json.NewEncoder(w)

	strID := r.URL.Query().Get("id")

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting a pallet \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		err = palletRepo.DeletePalletByID(id)

		if err != nil {
			log.Printf("Error deleting pallet: %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Error deleting pallet: %s", err.Error())})
			return
		}
	}

}

//PostPallet translates a POST request to a pallet creation
func (palletService *PalletService) PostPallet(w http.ResponseWriter, r *http.Request) {
	var palletRepo dao.PalletDao

	decoder := json.NewDecoder(r.Body)
	pallet := models.Pallet{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&pallet)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid pallet supplied"})
		return
	}

	pallet.ID = 0

	pal, err := palletRepo.SavePallet(pallet)

	encoder.Encode(pal)
}

//GetPallet translates a GET request to an pallet request from the db
func (palletService *PalletService) GetPallet(w http.ResponseWriter, r *http.Request) {

	var palletRepo dao.PalletDao
	var pieceRepo dao.PieceDao

	strID := r.URL.Query().Get("id")
	strPrintID := r.URL.Query().Get("print")

	encoder := json.NewEncoder(w)

	var pallet *models.Pallet

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting a pallet \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		pallet, err = palletRepo.FindPalletByID(id)

		if err != nil {
			log.Printf("Error retrieving pallet with id %d: %s \r\n", id, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

	} else if len(strPrintID)!=0 { 
		palletID, err := strconv.ParseInt(strPrintID, 10, 64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting a pallet \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}
		pallet, err := palletRepo.FindPalletByID(palletID)
		if err != nil {
			log.Printf("Error retrieving pallet with id %d: %s \r\n", palletID, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}
		pieces, err := pieceRepo.FindPiecesByPalletsID(palletID)
		
		w.Header().Add("Content-type","application/pdf")
		err = pdf.GeneratePalletPDF(*pallet,pieces,httputil.NewChunkedWriter(w))
		if err != nil {
			log.Printf("Error generating pdf: %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}
	} else {
		//return all pallets
		pallets, err := palletRepo.FindAllPallets()

		if err != nil {
			log.Printf("Error retrieving pallets %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

		encoder.Encode(pallets)
		return
	}

	if pallet != nil {
		encoder.Encode(*pallet)
	}

}
