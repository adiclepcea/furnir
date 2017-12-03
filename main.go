package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/adiclepcea/furnir/dao"
	"github.com/adiclepcea/furnir/service"
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

func operationsHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	switch r.Method {
	case http.MethodGet:
		service.PrintOperations(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(service.Error{Message: "Unknown method supplied"})
	}
}

func main() {

	database := flag.String("database", "furnir", "The database to use")
	dbUser := flag.String("dbuser", "furnir", "The database user")
	dbPassword := flag.String("dbpass", "Furnir123", "The database password")
	dbPort := flag.Int("dbport", 3306, "The port of the database")
	dbServer := flag.String("dbserver", "localhost", "The database server")
	port := flag.Int("port", 5000, "The port to start the server on")

	flag.Parse()

	if database != nil {
		dao.Database = *database
	}
	if dbUser != nil {
		dao.DbUser = *dbUser
	}
	if dbPassword != nil {
		dao.DbPassword = *dbPassword
	}
	if dbPort != nil {
		dao.DbPort = *dbPort
	}
	if dbServer != nil {
		dao.DbServer = *dbServer
	}

	defPort := 5000
	if port != nil {
		defPort = *port
	}

	fmt.Printf("port: %d\r\n", *port)

	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", http.StripPrefix("/", fs).ServeHTTP)
	essenceService = service.NewEssenceService(dao.EssenceDao{})
	http.HandleFunc("/pallet", palletHandler)
	http.HandleFunc("/essence", essenceHandler)
	http.HandleFunc("/piece", pieceHandler)
	http.HandleFunc("/transfer", transferHandler)
	http.HandleFunc("/operations", operationsHandler)
	fmt.Printf("Starting furnir server on port %d ...\r\n", defPort)
	http.ListenAndServe(fmt.Sprintf(":%d", defPort), nil)
}
