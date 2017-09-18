package main

import 
(
	"net/http"	
	"fmt"
	"log"
	"github.com/adiclepcea/furnir/models"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("OK"))
}

func main(){	
	pallet, err:=models.NewPallet()
	if err!=nil{
		log.Panic(err)
	}
	log.Println(pallet.ID)
	http.HandleFunc("/pallets",handler)
	http.ListenAndServe(":8080",nil)
	fmt.Println("llll")
}