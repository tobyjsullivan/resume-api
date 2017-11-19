package resolvers

import (
	"github.com/graphql-go/graphql"
	"errors"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
)

type city string

type cityData struct {
	Name string `json:"name"`
	CountryID string `json:"countryId"`
}

func (c city) getData() (*cityData, error) {
	log.Println("Fetching data for city:", string(c))
	resp, err := http.Get(fmt.Sprintf("http://places-db:3000/cities/%s", string(c)))
	if err != nil {
		log.Println("city.getData: ", err.Error())
		return nil, err
	}

	var respData struct {
		Result *cityData `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("city.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

var cityFields = graphql.Fields{
	"name": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			c, ok := p.Source.(city)
			if !ok {
				return nil, errors.New("Couldn't cast to City")
			}

			data, err := c.getData()
			if err != nil {
				return "", err
			}

			return data.Name, nil
		},
	},
	"country": &graphql.Field{
		Type: countryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			c, ok := p.Source.(city)
			if !ok {
				return nil, errors.New("Couldn't cast to City")
			}

			data, err := c.getData()
			if err != nil {
				return "", err
			}

			return country(data.CountryID), nil
		},
	},
}

var cityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "City",
		Fields: cityFields,
	},
)
