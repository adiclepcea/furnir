package service

import (
	"encoding/json"
	"fmt"
	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/models"
	"log"
	"net/http"
	"strconv"
)

//PieceService serves as a prototype for serving piece
//operations
type PieceService struct {
	pieceRepo dao.PieceDao
}

//NewPieceService returns a new instance of PieceService using the passed
//in dao.EssenceDao
func NewPieceService(pieceRepo dao.PieceDao) PieceService {
	ps := PieceService{pieceRepo: pieceRepo}
	return ps
}

//PutPiece translates a PUT request to an piece modification
func (pieceService *PieceService) PutPiece(w http.ResponseWriter, r *http.Request) {
	var pieceRepo dao.PieceDao

	decoder := json.NewDecoder(r.Body)
	piece := models.Piece{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&piece)

	sc, err := models.ScannedPiece{}.NewFromScan(piece.Barcode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid piece code supplied"})
		return
	}
	piece.Scanned = *sc

	if err != nil || piece.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid piece supplied"})
		return
	}

	ess, err := pieceRepo.SavePiece(piece)

	encoder.Encode(ess)
}

//DeletePiece translates a DELETE request to an piece deletion
func (pieceService *PieceService) DeletePiece(w http.ResponseWriter, r *http.Request) {
	var pieceRepo dao.PieceDao

	encoder := json.NewEncoder(w)

	strID := r.URL.Query().Get("id")

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting a piece \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		err = pieceRepo.DeletePieceByID(id)

		if err != nil {
			log.Printf("Error deleting piece: %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Error deleting piece: %s", err.Error())})
			return
		}
	}

}

//PostPiece translates a POST request to a piece creation
func (pieceService *PieceService) PostPiece(w http.ResponseWriter, r *http.Request) {
	var pieceRepo dao.PieceDao

	decoder := json.NewDecoder(r.Body)
	piece := models.Piece{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&piece)

	sc, err := models.ScannedPiece{}.NewFromScan(piece.Barcode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid piece code supplied"})
		return
	}
	piece.Scanned = *sc

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid piece supplied"})
		return
	}

	piece.ID = 0

	pc, err := pieceRepo.SavePiece(piece)

	encoder.Encode(pc)
}

//GetPiece translates a GET request to an piece request from the db
func (pieceService *PieceService) GetPiece(w http.ResponseWriter, r *http.Request) {

	var pieceRepo dao.PieceDao

	strID := r.URL.Query().Get("id")
	barcode := r.URL.Query().Get("barcode")
	palletID := r.URL.Query().Get("pallet_id")

	encoder := json.NewEncoder(w)

	var piece *models.Piece

	if len(strID) != 0 {
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			log.Printf("Invalid id (%s) received for getting a piece \r\n", strID)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(Error{Message: fmt.Sprintf("Invalid id: %s", strID)})
			return
		}

		piece, err = pieceRepo.FindPieceByID(id)

		if err != nil {
			log.Printf("Error retrieving piece with id %d: %s \r\n", id, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

	} else if len(barcode) != 0 {

		pieces, err := pieceRepo.FindPiecesByBarcode(barcode)

		if err != nil {
			log.Printf("Error retrieving piece with code %s: %s \r\n", barcode, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

		encoder.Encode(pieces)
		return

	} else if len(palletID) != 0 {
		pID, err := strconv.ParseInt(palletID, 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pieces, err := pieceRepo.FindPiecesByPalletsID(pID)

		if err != nil {
			log.Printf("Error retrieving piece from pallet %s: %s \r\n", palletID, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

		encoder.Encode(pieces)
		return
	} else {
		//return all essences
		pieces, err := pieceRepo.FindAllPieces()

		if err != nil {
			log.Printf("Error retrieving pieces %s \r\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(Error{Message: fmt.Sprintf("Internal error: %s", err.Error())})
			return
		}

		encoder.Encode(pieces)
		return
	}

	if piece != nil {
		encoder.Encode(*piece)
	}

}
