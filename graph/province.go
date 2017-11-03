package graph

import "github.com/satori/go.uuid"

type province struct {
	provinceId uuid.UUID
	name       string
}

func (p *province) id() uuid.UUID {
	return p.provinceId
}

func (p *province) applyRelation(nodes *map[uuid.UUID]node, r relation) {
	// No relationships
}
