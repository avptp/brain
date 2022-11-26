package resolvers

import (
	"github.com/avptp/brain/internal/generated/data"
)

type Resolver struct {
	data *data.Client
}

func NewResolver(data *data.Client) *Resolver {
	return &Resolver{
		data,
	}
}
