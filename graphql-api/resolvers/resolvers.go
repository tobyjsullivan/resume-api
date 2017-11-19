package resolvers

import (
	"github.com/graphql-go/graphql"
	"log"
)

const mePersonId = "40d8ed45-6977-47b5-92fa-7c4c4fe214c6"

func RootQueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me": &graphql.Field{
					Type: personType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						log.Println("me.Resolve: returning person", mePersonId)
						return person(mePersonId), nil
					},
				},
			},
		},
	)
}
