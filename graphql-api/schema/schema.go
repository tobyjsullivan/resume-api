package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/tobyjsullivan/resume-api/graphql-api/resolvers"
)

func NewSchema() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{Query: resolvers.RootQueryType()}

	return graphql.NewSchema(schemaConfig)
}
