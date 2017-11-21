package resolvers

import (
	"log"

	"github.com/graphql-go/graphql"
)

const mePersonId = "40d8ed45-6977-47b5-92fa-7c4c4fe214c6"

func buildObjectType(name string, fields func() graphql.Fields) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   name,
			Fields: graphql.FieldsThunk(fields),
		},
	)
}

func init() {
	cityType = buildObjectType("City", buildCityFields)
	companyType = buildObjectType("Company", buildCompanyFields)
	countryType = buildObjectType("Country", buildCountryFields)
	jobType = buildObjectType("Job", buildJobFields)
	personType = buildObjectType("Person", buildPersonFields)
	roleType = buildObjectType("Role", buildRoleFields)
}

func RootQueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me": &graphql.Field{
					Type: personType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						log.Println("me.Resolve: returning person", mePersonId)
						return &person{ID: mePersonId}, nil
					},
				},
			},
		},
	)
}
