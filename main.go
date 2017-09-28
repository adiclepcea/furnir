package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adiclepcea/furnir/models"
)

func palletHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPalletHandler(w, r)
		break
	case http.MethodPost:
		postPalletHandler(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func postPalletHandler(w http.ResponseWriter, r *http.Request) {
	pallet, err := models.NewPallet()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("A aparut o eroare la crearea unui nou palet"))
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(pallet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getPalletHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside getPalletHandler")
	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		log.Println("Inside getPallets")
		pallets, err := models.GetPallets()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encoder := json.NewEncoder(w)
		err = encoder.Encode(pallets)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		log.Println("Inside getPalletById " + id + " " + strconv.Itoa(len(id)))
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pallet, err := models.GetPalletByID(intid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encoder := json.NewEncoder(w)
		err = encoder.Encode(pallet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func main() {
	pallet, err := models.NewPallet()
	if err != nil {
		log.Panic(err)
	}
	log.Println(pallet.ID)
	http.HandleFunc("/pallet", palletHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("llll")
}
