package service

import (
	"log"
	"encoding/json"
	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/models"
	"net/http"
)

//TransferService serves as a prototype for serving transfer
//operations
type TransferService struct {
	pieceRepo  dao.PieceDao
	palletRepo dao.PalletDao
}

//NewTransferService returns a new instance of TransferService using the passed
//dao instances
func NewTransferService(palletRepo dao.PalletDao, pieceRepo dao.PieceDao) TransferService {
	es := TransferService{palletRepo: palletRepo, pieceRepo: pieceRepo}
	return es
}

//PostTransfer translates a POST request to an transfer creation
func (transferService *TransferService) PostTransfer(w http.ResponseWriter, r *http.Request) {

	var palletRepo dao.PalletDao
	var pieceRepo dao.PieceDao

	decoder := json.NewDecoder(r.Body)
	transfer := models.Transfer{}

	encoder := json.NewEncoder(w)

	err := decoder.Decode(&transfer)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Transfer invalid!"})
		return
	}

	pallet, err := palletRepo.FindPalletByID(transfer.SourcePalletID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Nu se poate obtine paletul sursa!"})
		return
	}

	if pallet == nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Paletul sursa nu exista!"})
		return
	}

	palletDest, err := palletRepo.FindPalletByID(transfer.DestPalletID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Paletul destinatie nu poate fi gasit!"})
		return
	}

	if palletDest == nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Paletul destinatie nu exista!"})
		return
	}

	_, err = models.ScannedPiece{}.NewFromScan(transfer.PieceBarcode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Pachet invalid!"})
		return
	}

	err = pieceRepo.TransferPieceByBarcode(transfer.PieceBarcode, transfer.SourcePalletID, transfer.DestPalletID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Nu se poate efectua transferul: "+err.Error()})
		return
	}

	encoder.Encode(transfer)
}
