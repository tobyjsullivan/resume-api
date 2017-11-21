package resolvers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type role struct {
	ID   string
	data *roleData
}

type roleData struct {
	JobID     string  `json:"jobId"`
	Title     string  `json:"title"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

func (r *role) getData() (*roleData, error) {
	if r.data == nil {
		log.Println("Fetching data for role:", r.ID)
		resp, err := http.Get(fmt.Sprintf("http://jobs-db:3000/roles/%s", r.ID))
		if err != nil {
			log.Println("role.getData: ", err.Error())
			return nil, err
		}

		var respData struct {
			Result *roleData `json:"result"`
			Error  string    `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
			log.Println("role.getData: ", err.Error())
			return nil, err
		}

		r.data = respData.Result
	}

	return r.data, nil
}

var roleType *graphql.Object

func buildRoleFields() graphql.Fields {
	return graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r, ok := p.Source.(*role)
				if !ok {
					return nil, errors.New("Couldn't cast to Role")
				}

				data, err := r.getData()
				if err != nil {
					return nil, err
				}

				return data.Title, nil
			},
		},
		"job": &graphql.Field{
			Type: jobType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r, ok := p.Source.(*role)
				if !ok {
					return nil, errors.New("Couldn't cast to Role")
				}

				data, err := r.getData()
				if err != nil {
					return nil, err
				}

				return &job{ID: data.JobID}, nil
			},
		},
		"startDate": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r, ok := p.Source.(*role)
				if !ok {
					return nil, errors.New("Couldn't cast to Role")
				}

				data, err := r.getData()
				if err != nil {
					return nil, err
				}

				return data.StartDate, nil
			},
		},
		"endDate": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				r, ok := p.Source.(*role)
				if !ok {
					return nil, errors.New("Couldn't cast to Role")
				}

				data, err := r.getData()
				if err != nil {
					return nil, err
				}

				return data.EndDate, nil
			},
		},
	}
}
