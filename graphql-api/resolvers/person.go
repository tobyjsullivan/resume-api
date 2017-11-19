package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type person string

type personData struct {
	GivenNames []string `json:"givenNames"`
	Surname    string   `json:"surname"`
	CurrentCityID string `json:"currentCityId"`
	HometownCityID string `json:"hometownCityId"`
}

func (p person) getData() (*personData, error) {
	log.Println("Fetching data for person:", string(p))
	resp, err := http.Get(fmt.Sprintf("http://people-db:3000/people/%s", string(p)))
	if err != nil {
		log.Println("person.getData: ", err.Error())
		return nil, err
	}

	var respData struct {
		Result *personData `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("person.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"firstName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					data, err := person.getData()
					if err != nil {
						return []string{}, err
					}

					var firstName string
					if len(data.GivenNames) > 0 {
						firstName = data.GivenNames[0]
					}

					return firstName, nil
				},
			},
			"surname": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					data, err := person.getData()
					if err != nil {
						return "", err
					}

					return data.Surname, nil
				},
			},
			"givenNames": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					data, err := person.getData()
					if err != nil {
						return []string{}, err
					}

					return data.GivenNames, nil
				},
			},
			"currentCity": &graphql.Field{
				Type: cityType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					data, err := person.getData()
					if err != nil {
						return nil, err
					}

					return city(data.CurrentCityID), nil
				},
			},
			"hometown": &graphql.Field{
				Type: cityType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					data, err := person.getData()
					if err != nil {
						return nil, err
					}

					return city(data.HometownCityID), nil
				},
			},
		},
	},
)
