package main

import (
	"auth_mysibsau/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func auth(w http.ResponseWriter, _ *http.Request) {
	tmpUser := models.User{Id: "1234", Name: "1235", Group: models.Group{Id: "1234", Name: "1235"}}
	err := json.NewEncoder(w).Encode(tmpUser)
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth", auth)
	log.Fatal(http.ListenAndServe(":8080", router))

}