package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
	"errors"
)

func resolveMe(db *data.Database) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return db.Me, nil
	}
}



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
		},
	},
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
		},
	},
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

					return person.Jobs, nil
				},
			},
		},
	},
)


func queryType(db *data.Database) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me": &graphql.Field{
					Type: personType,
					Resolve: resolveMe(db),
				},
			},
		},
	)
}

func NewSchema(db *data.Database) (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{Query: queryType(db)}

	return graphql.NewSchema(schemaConfig)
}
