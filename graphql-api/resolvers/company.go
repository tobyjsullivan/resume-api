package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type company struct {
	ID string
	data *companyData
}

type companyData struct {
	OfficialName string `json:"officialName"`
	CommonName string `json:"commonName"`
	CityID string `json:"cityId"`
}

func (c *company) getData() (*companyData, error) {
	if c.data == nil {
		log.Println("Fetching data for company:", c.ID)
		resp, err := http.Get(fmt.Sprintf("http://companies-db:3000/companies/%s", c.ID))
		if err != nil {
			log.Println("company.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *companyData `json:"result"`
			Error  string       `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("company.getData: ", err.Error())
			return nil, err
		}

		c.data = respData.Result
	}

	return c.data, nil
}

var companyType *graphql.Object

func buildCompanyFields() graphql.Fields {
	return graphql.Fields{
		"commonName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c, ok := p.Source.(*company)
				if !ok {
					return nil, errors.New("Couldn't cast to City")
				}

				data, err := c.getData()
				if err != nil {
					return nil, err
				}

				return data.CommonName, nil
			},
		},
		"officialName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c, ok := p.Source.(*company)
				if !ok {
					return nil, errors.New("Couldn't cast to City")
				}

				data, err := c.getData()
				if err != nil {
					return nil, err
				}

				return data.OfficialName, nil
			},
		},
		"city": &graphql.Field{
			Type: cityType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				c, ok := p.Source.(*company)
				if !ok {
					return nil, errors.New("Couldn't cast to City")
				}

				data, err := c.getData()
				if err != nil {
					return nil, err
				}

				return &city{ID: data.CityID}, nil
			},
		},
	}
}
