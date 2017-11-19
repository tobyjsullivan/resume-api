package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type country string

type countryData struct {
	CommonName string `json:"commonName"`
}

func (c country) getData() (*countryData, error) {
	log.Println("Fetching data for country:", string(c))
	resp, err := http.Get(fmt.Sprintf("http://places-db:3000/countries/%s", string(c)))
	if err != nil {
		log.Println("country.getData: ", err.Error())
		return nil, err
	}

	var respData struct {
		Result *countryData `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("country.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

var countryFields = graphql.Fields{
	"commonName": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			c, ok := p.Source.(country)
			if !ok {
				return nil, errors.New("Couldn't cast to Country")
			}

			data, err := c.getData()
			if err != nil {
				return nil, err
			}

			return data.CommonName, nil
		},
	},
}

var countryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Country",
		Fields: countryFields,
	},
)