package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"encoding/json"
	"log"
)

func main() {
	r := buildRoutes()

	n := negroni.New()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)
}

func buildRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", statusHandler).Methods("GET")
	r.HandleFunc("/people/{person-id}", peopleHandler).Methods("GET")

	return r
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string `json:"status"`
		Service string `json:"service"`
	}{
		Status: "ok",
		Service: "people-db",
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["person-id"]


}
