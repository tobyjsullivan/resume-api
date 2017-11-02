package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

var jobType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Job",
		Fields: graphql.Fields{
			"company": &graphql.Field{
				Type: companyType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					job, ok := p.Source.(*data.Job)
					if !ok {
						return nil, errors.New("Couldn't cast to Job")
					}

					return job.Company, nil
				},
			},
			"roles": &graphql.Field{
				Type: graphql.NewList(roleType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					job, ok := p.Source.(*data.Job)
					if !ok {
						return nil, errors.New("Couldn't cast to Job")
					}

					return job.Roles, nil
				},
			},
		},
	},
)
