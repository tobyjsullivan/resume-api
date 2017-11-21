package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type city struct {
	ID   string
	data *cityData
}

type cityData struct {
	Name      string `json:"name"`
	CountryID string `json:"countryId"`
}

func (c *city) getData() (*cityData, error) {
	if c.data == nil {
		log.Println("Fetching data for city:", c.ID)
		resp, err := http.Get(fmt.Sprintf("http://places-db:3000/cities/%s", c.ID))
		if err != nil {
			log.Println("city.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *cityData `json:"result"`
			Error  string    `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("city.getData: ", err.Error())
			return nil, err
		}

		c.data = respData.Result
	}

	return c.data, nil
}

var cityType *graphql.Object

func buildCityFields() graphql.Fields {
	return graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c, ok := p.Source.(*city)
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
				c, ok := p.Source.(*city)
				if !ok {
					return nil, errors.New("Couldn't cast to City")
				}

				data, err := c.getData()
				if err != nil {
					return "", err
				}

				return &country{ID: data.CountryID}, nil
			},
		},
	}
}
