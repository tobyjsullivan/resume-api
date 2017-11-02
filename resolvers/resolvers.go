package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/data"
)

const dateFmt = "2006-01-02"

func rootQueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me": &graphql.Field{
					Type: personType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return data.Me(), nil
					},
				},
			},
		},
	)
}

func NewSchema() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{Query: rootQueryType()}

	return graphql.NewSchema(schemaConfig)
}
