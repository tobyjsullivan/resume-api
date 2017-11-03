package graph

import "github.com/satori/go.uuid"

type node interface {
	ID() uuid.UUID
	ApplyRelation(nodes *map[uuid.UUID]node, r relation)
}
