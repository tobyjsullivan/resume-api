package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

var personType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"firstName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*data.Person)
					if !ok {
						return nil, errors.New("Couldn't cast to person")
					}

					return person.FirstName, nil
				},
			},
			"middleName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*data.Person)
					if !ok {
						return nil, errors.New("Couldn't cast to person")
					}

					return person.MiddleName, nil
				},
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*data.Person)
					if !ok {
						return nil, errors.New("Couldn't cast to person")
					}

					return person.LastName, nil
				},
			},
			"website": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*data.Person)
					if !ok {
						return nil, errors.New("Couldn't cast to person")
					}

					return person.Website.String(), nil
				},
			},
			"jobs": &graphql.Field{
				Type: graphql.NewList(jobType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					person, ok := p.Source.(*data.Person)
					if !ok {
						return nil, errors.New("Couldn't cast to person")
					}

					return person.JobHistory, nil
				},
			},
		},
	},
)

