package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type country struct {
	ID string
	data *countryData
}

type countryData struct {
	CommonName string `json:"commonName"`
}

func (c *country) getData() (*countryData, error) {
	if c.data == nil {
		log.Println("Fetching data for country:", c.ID)
		resp, err := http.Get(fmt.Sprintf("http://places-db:3000/countries/%s", c.ID))
		if err != nil {
			log.Println("country.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *countryData `json:"result"`
			Error  string       `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("country.getData: ", err.Error())
			return nil, err
		}

		c.data = respData.Result
	}

	return c.data, nil
}

var countryType *graphql.Object

func buildCountryFields() graphql.Fields {
	return graphql.Fields{
		"commonName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c, ok := p.Source.(*country)
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
}
