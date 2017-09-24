package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adiclepcea/furnir/models"
)

func palletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getPalletHandler(w, r)
		return
	}
}

func getPalletHandler(w http.ResponseWriter, r *http.Request) {

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
