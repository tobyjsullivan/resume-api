package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

var roleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Role",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					job, ok := p.Source.(*data.Role)
					if !ok {
						return nil, errors.New("Couldn't cast to Role")
					}

					return job.Title, nil
				},
			},
			"startDate": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					job, ok := p.Source.(*data.Role)
					if !ok {
						return nil, errors.New("Couldn't cast to Role")
					}

					if job.StartDate == nil {
						return nil, nil
					}

					return job.StartDate.Format(dateFmt), nil
				},
			},
			"endDate": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					job, ok := p.Source.(*data.Role)
					if !ok {
						return nil, errors.New("Couldn't cast to Role")
					}

					if job.EndDate == nil {
						return nil, nil
					}

					return job.EndDate.Format(dateFmt), nil
				},
			},
		},
	},
)

