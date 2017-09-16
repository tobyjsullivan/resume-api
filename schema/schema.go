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
