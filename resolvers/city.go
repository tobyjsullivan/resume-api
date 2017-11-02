package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

var cityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "City",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					city, ok := p.Source.(*data.City)
					if !ok {
						return nil, errors.New("Couldn't cast to City")
					}

					return city.Name, nil
				},
			},
		},
	},
)
