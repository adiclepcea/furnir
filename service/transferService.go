package service

import (
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
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Invalid transfer supplied"})
		return
	}

	pallet, err := palletRepo.FindPalletByID(transfer.SourcePalletID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Could not obtain source pallet"})
		return
	}

	if pallet == nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Source pallet does not exist"})
		return
	}

	palletDest, err := palletRepo.FindPalletByID(transfer.DestPalletID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Could not obtain destination pallet"})
		return
	}

	if palletDest == nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&Error{Message: "Destination pallet does not exist"})
		return
	}

	err = pieceRepo.TransferPieceByBarcode(transfer.PieceBarcode, transfer.SourcePalletID, transfer.DestPalletID)
	if err != nil {

	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(&Error{Message: "Could not transfer piece"})
		return
	}

	encoder.Encode(transfer)
}
