package main

import (
	"net/http"
	"os"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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
	r.HandleFunc("/", statusHandler).Methods(http.MethodGet)
	r.HandleFunc("/people", peopleIndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/people/{id}", personHandler).Methods(http.MethodGet)

	return r
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status  string `json:"status"`
		Service string `json:"service"`
	}{
		Status:  "ok",
		Service: "people-db",
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func peopleIndexHandler(w http.ResponseWriter, r *http.Request) {

}

func personHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId := vars["id"]

	p := find(personId)

	if p == nil {
		http.Error(w, "No person found.", http.StatusNotFound)
		return
	}

	res := &response{
		Result: p,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}
