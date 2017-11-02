package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

var companyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, ok := p.Source.(*data.Company)
					if !ok {
						return nil, errors.New("Couldn't cast to Company")
					}

					return company.Name, nil
				},
			},
			"city": &graphql.Field{
				Type: cityType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					company, ok := p.Source.(*data.Company)
					if !ok {
						return nil, errors.New("Couldn't cast to Company")
					}

					return company.City, nil
				},
			},
		},
	},
)

