package graph

import "github.com/satori/go.uuid"

type city struct {
	cityId   uuid.UUID
	name     string
	province *province
}

func (c *city) id() uuid.UUID {
	return c.cityId
}

func (c *city) applyRelation(nodes *map[uuid.UUID]node, r relation) {
	switch rel := r.(type) {
	case *cityInProvinceRelation:
		p := (*nodes)[rel.provinceId].(*province)
		c.province = p
	}
}

type cityInProvinceRelation struct {
	cityId     uuid.UUID
	provinceId uuid.UUID
}

func (r *cityInProvinceRelation) from() uuid.UUID {
	return r.cityId
}
