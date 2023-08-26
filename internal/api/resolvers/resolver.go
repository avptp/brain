package resolvers

import (
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/messaging"
)

type Resolver struct {
	data      *data.Client
	messenger messaging.Messenger
}

func NewResolver(data *data.Client, messenger messaging.Messenger) *Resolver {
	return &Resolver{
		data,
		messenger,
	}
}
