package graph

import "github.com/satori/go.uuid"

type relation interface {
	From() uuid.UUID
}

