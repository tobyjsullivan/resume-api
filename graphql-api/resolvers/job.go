package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type job struct {
	ID       string
	data     *jobData
	roleData *[]*role
}

type jobData struct {
	EmployeePersonId  string  `json:"employeePersonId"`
	EmployerCompanyId string  `json:"employerCompanyId"`
	LocationCityId    string  `json:"locationCityId"`
	Remote            bool    `json:"remote"`
	StartDate         *string `json:"startDate"`
	EndDate           *string `json:"endDate"`
}

func (j *job) getData() (*jobData, error) {
	if j.data == nil {
		log.Println("Fetching data for job:", j.ID)
		resp, err := http.Get(fmt.Sprintf("http://jobs-db:3000/jobs/%s", j.ID))
		if err != nil {
			log.Println("job.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *jobData `json:"result"`
			Error  string   `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("job.getData: ", err.Error())
			return nil, err
		}

		j.data = respData.Result
	}

	return j.data, nil
}

func (j *job) getRoles() ([]*role, error) {
	if j.roleData == nil {
		log.Println("Fetching roles for job:", j.ID)
		req, err := http.NewRequest(http.MethodGet, "http://jobs-db:3000/roles", nil)
		if err != nil {
			log.Println("job.getRoles: ", err.Error())
			return []*role{}, err
		}

		q := req.URL.Query()
		q.Set("job-id", j.ID)
		req.URL.RawQuery = q.Encode()

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("job.getRoles: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result []struct {
				ID string `json:"id"`
			} `json:"result"`
			Error string `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("job.getRoles: ", err.Error())
			return nil, err
		}

		out := []*role{}
		for _, roleData := range respData.Result {
			out = append(out, &role{ID: roleData.ID})
		}

		j.roleData = &out
	}

	return *j.roleData, nil
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
		"startDate": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return data.StartDate, nil
			},
		},
		"endDate": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				data, err := j.getData()
				if err != nil {
					return nil, err
				}

				return data.EndDate, nil
			},
		},
		"roles": &graphql.Field{
			Type: graphql.NewList(roleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				j, ok := p.Source.(*job)
				if !ok {
					return nil, errors.New("Couldn't cast to Job")
				}

				roles, err := j.getRoles()
				if err != nil {
					return nil, err
				}

				return roles, nil
			},
		},
	}
}
