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
	r.HandleFunc("/jobs", jobIndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/jobs/{id}", jobHandler).Methods(http.MethodGet)

	return r
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Status string `json:"status"`
		Service string `json:"service"`
	}{
		Status: "ok",
		Service: "jobs-db",
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func jobIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("jobIndexHandler: request received.")
	q := r.URL.Query()

	personId := q.Get("person-id")
	if personId == "" {
		respondWithError(w, "param person-id must be provided", http.StatusBadRequest)
		return
	}

	result := findByPersonId(personId)

	res := &response{
		Result: result,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
	}

	//respond(w, jobs)
}

func jobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("jobHandler: request received.")
	vars := mux.Vars(r)
	id := vars["id"]

	c := find(id)

	if c == nil {
		log.Println("jobHandler: no company found")
		http.Error(w, "No company found.", http.StatusNotFound)
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
