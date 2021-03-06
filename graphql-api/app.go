package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/graphql-api/schema"
	"github.com/urfave/negroni"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	graphSchema, err := schema.NewSchema()
	if err != nil {
		log.Fatalln("failed to create new resolvers", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request to /graphql")
		var req struct {
			Query     string      `json:"query"`
			Variables interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := executeQuery(req.Query, graphSchema)
		json.NewEncoder(w).Encode(result)
		return
	}).Methods("POST")

	n := negroni.New()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Now server is running on port %s\n", port)
	n.Run(":" + port)
}
