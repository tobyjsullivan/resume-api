package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type person struct {
	ID      string
	data    *personData
	jobData *[]*job
}

type personData struct {
	GivenNames     []string `json:"givenNames"`
	Surname        string   `json:"surname"`
	CurrentCityID  string   `json:"currentCityId"`
	HometownCityID string   `json:"hometownCityId"`
}

func (p *person) getData() (*personData, error) {
	if p.data == nil {
		log.Println("Fetching data for person:", p.ID)
		resp, err := http.Get(fmt.Sprintf("http://people-db:3000/people/%s", p.ID))
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

		p.data = respData.Result
	}

	return p.data, nil
}

func (p *person) getJobs() ([]*job, error) {
	if p.jobData == nil {
		log.Println("Fetching jobs for person:", p.ID)
		req, err := http.NewRequest(http.MethodGet, "http://jobs-db:3000/jobs", nil)
		if err != nil {
			log.Println("person.getJobs: ", err.Error())
			return []*job{}, err
		}

		q := req.URL.Query()
		q.Set("person-id", p.ID)
		req.URL.RawQuery = q.Encode()

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("person.getJobs: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result []struct {
				ID string `json:"id"`
			} `json:"result"`
			Error string `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("person.getData: ", err.Error())
			return nil, err
		}

		out := []*job{}
		for _, jData := range respData.Result {
			out = append(out, &job{ID: jData.ID})
		}

		p.jobData = &out
	}

	return *p.jobData, nil
}

var personType *graphql.Object

func buildPersonFields() graphql.Fields {
	return graphql.Fields{
		"firstName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				data, err := person.getData()
				if err != nil {
					return nil, err
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
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				data, err := person.getData()
				if err != nil {
					return nil, err
				}

				return data.Surname, nil
			},
		},
		"givenNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				data, err := person.getData()
				if err != nil {
					return nil, err
				}

				return data.GivenNames, nil
			},
		},
		"currentCity": &graphql.Field{
			Type: cityType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				data, err := person.getData()
				if err != nil {
					return nil, err
				}

				return &city{ID: data.CurrentCityID}, nil
			},
		},
		"hometown": &graphql.Field{
			Type: cityType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				data, err := person.getData()
				if err != nil {
					return nil, err
				}

				return &city{ID: data.HometownCityID}, nil
			},
		},
		"jobs": &graphql.Field{
			Type: graphql.NewList(jobType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person, ok := p.Source.(*person)
				if !ok {
					log.Println("personType.Resolve: case failes")
					return nil, errors.New("Couldn't cast to person")
				}

				jobs, err := person.getJobs()
				if err != nil {
					return nil, err
				}

				return jobs, nil
			},
		},
	}
}
