package resolvers

import (
	"github.com/graphql-go/graphql"
	"errors"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

type person struct {
	id string
}

type personData struct {
	GivenNames []string `json:"givenNames"`
	Surname string `json:"surname"`
}

func (p *person) getData() (*personData, error) {
	log.Println("Fetching data for person:", p.id)
	resp, err := http.Get(fmt.Sprintf("http://people-db:3000/people/%s", p.id))
	if err != nil {
		log.Println("person.getData: ", err.Error())
		return nil, err
	}

	var respData struct{
		Result *personData `json:"result"`
		Error string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("person.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

func (p *person) firstName() (string, error) {
	data, err := p.getData()
	if err != nil {
		return "", err
	}

	log.Println("firstName: GivenNames:", data.GivenNames)

	var firstName string
	if len(data.GivenNames) > 0 {
		firstName = data.GivenNames[0]
	}

	return firstName, nil
}

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"firstName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*person)
					if !ok {
						log.Println("personType.Resolve: case failes")
						return nil, errors.New("Couldn't cast to person")
					}

					return person.firstName()
				},
			},
		},
	},
)

