package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type company string

type companyData struct {
	OfficialName string `json:"officialName"`
	CommonName string `json:"commonName"`
	CityID string `json:"cityId"`
}

func (c company) getData() (*companyData, error) {
	log.Println("Fetching data for company:", string(c))
	resp, err := http.Get(fmt.Sprintf("http://companies-db:3000/companies/%s", string(c)))
	if err != nil {
		log.Println("company.getData: ", err.Error())
		return nil, err
	}

	var respData struct {
		Result *companyData `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("company.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

var companyFields = graphql.Fields{
	"commonName": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			c, ok := p.Source.(company)
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
	"city": &graphql.Field{
		Type: countryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			c, ok := p.Source.(company)
			if !ok {
				return nil, errors.New("Couldn't cast to City")
			}

			data, err := c.getData()
			if err != nil {
				return nil, err
			}

			return city(data.CityID), nil
		},
	},
}

var companyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: companyFields,
	},
)