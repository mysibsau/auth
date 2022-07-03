package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"src/odoo"
	"src/shemas"
)

func auth(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if recover() != nil {
			http.Error(w, "Что-то пошло не так", http.StatusInternalServerError)
		}
	}()

	var a shemas.Auth
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := odoo.Auth(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/auth", auth).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
