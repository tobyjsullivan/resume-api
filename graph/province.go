package graph

import "github.com/satori/go.uuid"

type province struct {
	provinceId uuid.UUID
	name       string
}

func (p *province) ID() uuid.UUID {
	return p.provinceId
}

func (p *province) ApplyRelation(nodes *map[uuid.UUID]node, r relation) {
	// No relationships
}
