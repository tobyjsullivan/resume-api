package graph

import "github.com/satori/go.uuid"

type city struct {
	cityId   uuid.UUID
	name     string
	province *province
}

func (c *city) ID() uuid.UUID {
	return c.cityId
}

func (c *city) ApplyRelation(nodes *map[uuid.UUID]node, r relation) {
	switch rel := r.(type) {
	case *cityInProvinceRelation:
		p := (*nodes)[rel.provinceId].(*province)
		c.province = p
	}
}

type cityInProvinceRelation struct {
	cityId uuid.UUID
	provinceId uuid.UUID
}

func (r *cityInProvinceRelation) From() uuid.UUID {
	return r.cityId
}
