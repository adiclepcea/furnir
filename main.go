package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adiclepcea/furnir/service"
	"github.com/adiclepcea/furnir/dao"
)

var essenceService service.EssenceService
var palletService service.PalletService

func essenceHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	if r.Method == http.MethodGet {
		essenceService.GetEssence(w, r)
		return
	}else if r.Method == http.MethodPost {
		essenceService.PostEssence(w, r)
		return
	}else if r.Method == http.MethodPut {
		essenceService.PutEssence(w, r)
		return
	}else if r.Method == http.MethodDelete {
		essenceService.DeleteEssence(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	encoder.Encode(service.Error{Message:"Unknown method supplied"})
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
		palletService.DeletePallet(w,r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(service.Error{Message:"Unknown method supplied"})
	}
}

func main() {
	//pallet, err := models.NewPallet()
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.Println(pallet.ID)
	essenceService = service.NewEssenceService(dao.EssenceDao{})
	http.HandleFunc("/pallet", palletHandler)
	http.HandleFunc("/essence", essenceHandler)
	http.ListenAndServe(":5000", nil)
	fmt.Println("llll")
}
