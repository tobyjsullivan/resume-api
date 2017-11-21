package resolvers

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"errors"
)

type job struct {
	ID string
	data *jobData
}

type jobData struct {
	EmployeePersonId string `json:"employeePersonId"`
	EmployerCompanyId string `json:"employerCompanyId"`
	LocationCityId string `json:"locationCityId"`
	Remote bool `json:"remote"`
}

func (j *job) getData() (*jobData, error) {
	if j.data == nil {
		log.Println("Fetching data for job:", j.ID)
		resp, err := http.Get(fmt.Sprintf("http://jobs-db:3000/jobs/%s", j.ID))
		if err != nil {
			log.Println("city.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *jobData `json:"result"`
			Error  string   `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("city.getData: ", err.Error())
			return nil, err
		}

		j.data = respData.Result
	}

	return j.data, nil
}

var jobType *graphql.Object

func buildJobFields() graphql.Fields {
	return graphql.Fields{
		"employee": &graphql.Field{
			Type: personType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return &person{ID: data.EmployeePersonId}, nil
			},
		},
		"employer": &graphql.Field{
			Type: companyType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return &company{ID: data.EmployerCompanyId}, nil
			},
		},
		"location": &graphql.Field{
			Type: cityType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return &city{ID: data.LocationCityId}, nil
			},
		},
		"remote": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
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
