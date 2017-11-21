package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type job string

type jobData struct {
	EmployeePersonId string `json:"employeePersonId"`
	EmployerCompanyId string `json:"employerCompanyId"`
	LocationCityId string `json:"locationCityId"`
	Remote bool `json:"remote"`
}

func (j job) getData() (*jobData, error) {
	log.Println("Fetching data for job:", string(j))
	resp, err := http.Get(fmt.Sprintf("http://jobs-db:3000/jobs/%s", string(j)))
	if err != nil {
		log.Println("city.getData: ", err.Error())
		return nil, err
	}

	var respData struct {
		Result *jobData `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Println("city.getData: ", err.Error())
		return nil, err
	}

	return respData.Result, nil
}

var jobType *graphql.Object

func buildJobFields() graphql.Fields {
	return graphql.Fields{
		"employee": &graphql.Field{
			Type: personType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return person(data.EmployeePersonId), nil
			},
		},
		"employer": &graphql.Field{
			Type: companyType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return company(data.EmployerCompanyId), nil
			},
		},
		"location": &graphql.Field{
			Type: cityType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return city(data.LocationCityId), nil
			},
		},
		"remote": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return data.Remote, nil
			},
		},
	}
}
