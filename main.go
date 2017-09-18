package main

import 
(
	"net/http"	
	"fmt"
	"github.com/adiclepcea/furnir/models"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("OK"))
}

func main(){
	models.InitDB()
	http.HandleFunc("/pallets",handler)
	http.ListenAndServe(":8080",nil)
	fmt.Println("llll")
}