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
	r.HandleFunc("/", statusHandler).Methods(http.MethodGet)
	r.HandleFunc("/countries/{id}", countryHandler).Methods(http.MethodGet)
	r.HandleFunc("/cities/{id}", cityHandler).Methods(http.MethodGet)

	return r
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string `json:"status"`
		Service string `json:"service"`
	}{
		Status: "ok",
		Service: "places-db",
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func countryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := findCountry(id)

	if c == nil {
		respondWithError(w, "No country found", http.StatusNotFound)
		return
	}

	respond(w, c)
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := findCity(id)

	if c == nil {
		respondWithError(w, "No city found", http.StatusNotFound)
		return
	}

	respond(w, c)
}

func respond(w http.ResponseWriter, result interface{}) {
	res := &response{
		Result: result,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
	}
}

func respondWithError(w http.ResponseWriter, err string, code int) {
	res := &response{
		Error: err,
	}

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Panic(err)
	}
}

type response struct {
	Result interface{} `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}
