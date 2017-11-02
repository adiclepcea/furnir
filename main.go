package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/service"
	"github.com/adiclepcea/furnir/pdf"
	"github.com/adiclepcea/furnir/models"
)

var essenceService service.EssenceService
var palletService service.PalletService
var pieceService service.PieceService
var transferService service.TransferService

func essenceHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	if r.Method == http.MethodGet {
		essenceService.GetEssence(w, r)
		return
	} else if r.Method == http.MethodPost {
		essenceService.PostEssence(w, r)
		return
	} else if r.Method == http.MethodPut {
		essenceService.PutEssence(w, r)
		return
	} else if r.Method == http.MethodDelete {
		essenceService.DeleteEssence(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	encoder.Encode(service.Error{Message: "Unknown method supplied"})
}

func palletHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	switch r.Method {
	case http.MethodGet:
		palletService.GetPallet(w, r)
		break
	case http.MethodPost:
		palletService.PostPallet(w, r)
		break
	case http.MethodPut:
		palletService.PutPallet(w, r)
		break
	case http.MethodDelete:
		palletService.DeletePallet(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(service.Error{Message: "Unknown method supplied"})
	}
}

func pieceHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	switch r.Method {
	case http.MethodGet:
		pieceService.GetPiece(w, r)
		break
	case http.MethodPost:
		pieceService.PostPiece(w, r)
		break
	case http.MethodPut:
		pieceService.PutPiece(w, r)
		break
	case http.MethodDelete:
		pieceService.DeletePiece(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(service.Error{Message: "Unknown method supplied"})
	}
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	switch r.Method {
	case http.MethodPost:
		transferService.PostTransfer(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(service.Error{Message: "Unknown method supplied"})
	}
}

func main() {
	sp1,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp2,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp3,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	sp4,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp5,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp6,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	sp7,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp8,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp9,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	sp10,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp11,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp12,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	sp13,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp14,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp15,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	sp16,_ :=  models.ScannedPiece{}.NewFromScan("1100214400000342025514")
	sp17,_ :=  models.ScannedPiece{}.NewFromScan("1200244400000342025722")
	sp18,_ :=  models.ScannedPiece{}.NewFromScan("2200154400000343325722")
	pieces := []models.Piece{
		models.Piece{ID:1,Barcode:"1100214400000342025514", Scanned: *sp1 },
		models.Piece{ID:2,Barcode:"1200244400000342025722", Scanned: *sp2},
		models.Piece{ID:3,Barcode:"2200154400000343325722", Scanned: *sp3},
		models.Piece{ID:4,Barcode:"1100214400000342025514", Scanned: *sp4 },
		models.Piece{ID:5,Barcode:"1200244400000342025722", Scanned: *sp5},
		models.Piece{ID:6,Barcode:"2200154400000343325722", Scanned: *sp6},
		models.Piece{ID:7,Barcode:"1100214400000342025514", Scanned: *sp7 },
		models.Piece{ID:8,Barcode:"1200244400000342025722", Scanned: *sp8},
		models.Piece{ID:9,Barcode:"2200154400000343325722", Scanned: *sp9},
		models.Piece{ID:10,Barcode:"1100214400000342025514", Scanned: *sp10 },
		models.Piece{ID:11,Barcode:"1200244400000342025722", Scanned: *sp11},
		models.Piece{ID:12,Barcode:"2200154400000343325722", Scanned: *sp12},
		models.Piece{ID:13,Barcode:"1100214400000342025514", Scanned: *sp13 },
		models.Piece{ID:14,Barcode:"1200244400000342025722", Scanned: *sp14},
		models.Piece{ID:15,Barcode:"2200154400000343325722", Scanned: *sp15},
		models.Piece{ID:16,Barcode:"1100214400000342025514", Scanned: *sp16 },
		models.Piece{ID:17,Barcode:"1200244400000342025722", Scanned: *sp17},
		models.Piece{ID:18,Barcode:"2200154400000343325722", Scanned: *sp18},
	}
	pdf.GeneratePalletPDF(models.Pallet{ID:10,Essence:models.Essence{Name:"Stejar"}},pieces)
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", http.StripPrefix("/", fs).ServeHTTP)
	essenceService = service.NewEssenceService(dao.EssenceDao{})
	http.HandleFunc("/pallet", palletHandler)
	http.HandleFunc("/essence", essenceHandler)
	http.HandleFunc("/piece", pieceHandler)
	http.HandleFunc("/transfer", transferHandler)
	fmt.Println("Starting furnir server on port 5000 ...")
	http.ListenAndServe(":5000", nil)
	fmt.Println("llll")
}
