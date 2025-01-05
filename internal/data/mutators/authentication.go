package mutators

import (
	"context"

	"entgo.io/ent"
	"github.com/avptp/brain/internal/crypto"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/hook"
)

func AuthenticationToken(next ent.Mutator) ent.Mutator {
	return hook.AuthenticationFunc(func(ctx context.Context, m *data.AuthenticationMutation) (ent.Value, error) {
		if _, ok := m.Token(); !ok {
			token, err := crypto.RandomBytes(64)

			if err != nil {
				return nil, err
			}

			m.SetToken(token)
		}

		return next.Mutate(ctx, m)
	})
}
