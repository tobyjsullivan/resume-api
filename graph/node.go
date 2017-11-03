package graph

import "github.com/satori/go.uuid"

type node interface {
	id() uuid.UUID
	applyRelation(nodes *map[uuid.UUID]node, r relation)
}
