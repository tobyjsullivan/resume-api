package graph

import (
	"github.com/satori/go.uuid"
	"fmt"
)

var (
	vancouverId = uuid.NewV4()
	victoriaId = uuid.NewV4()
	sanFranciscoId = uuid.NewV4()
	britishColumbiaId = uuid.NewV4()
	californiaId = uuid.NewV4()
)

var log = []interface{}{
	&defineNode{node: &city{cityId:vancouverId, name: "Vancouver"}},
	&defineNode{node: &city{cityId:victoriaId, name: "Victoria"}},
	&defineNode{node: &city{cityId:sanFranciscoId, name: "San Francisco"}},
	&defineNode{node: &province{provinceId: britishColumbiaId, name: "British Columbia"}},
	&defineNode{node: &province{provinceId: californiaId, name: "California"}},
	&defineRelation{relation: &cityInProvinceRelation{cityId: vancouverId, provinceId: britishColumbiaId}},
	&defineRelation{relation: &cityInProvinceRelation{cityId: victoriaId, provinceId: britishColumbiaId}},
	&defineRelation{relation: &cityInProvinceRelation{cityId: sanFranciscoId, provinceId: californiaId}},
}

func ProcessLog() {
	nodes := make(map[uuid.UUID]node)

	for _, event := range log {
		switch e := event.(type) {
		case *defineNode:
			n := e.Node()
			nodes[n.ID()] = n
		case *defineRelation:
			r := e.Relationship()
			f := nodes[r.From()]
			f.ApplyRelation(&nodes, r)
		}
	}

	vic := nodes[victoriaId].(*city)
	fmt.Printf("city: %v; province: %v\n", vic.name, vic.province.name)
	van := nodes[vancouverId].(*city)
	fmt.Printf("city: %v; province: %v\n", van.name, van.province.name)
	sf := nodes[sanFranciscoId].(*city)
	fmt.Printf("city: %v; province: %v\n", sf.name, sf.province.name)
}

type defineNode struct {
	node node
}

func (e *defineNode) Node() node {
	return e.node
}

type defineRelation struct {
	relation relation
}

func (e *defineRelation) Relationship() relation {
	return e.relation
}
