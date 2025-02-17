package mutators

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/alexedwards/argon2id"
	"github.com/avptp/brain/internal/generated/data"
	"github.com/avptp/brain/internal/generated/data/hook"
	"golang.org/x/text/language"
)

func PersonEmail(next ent.Mutator) ent.Mutator {
	return hook.PersonFunc(func(ctx context.Context, m *data.PersonMutation) (ent.Value, error) {
		if _, ok := m.Email(); ok {
			m.ClearEmailVerifiedAt()
		}

		return next.Mutate(ctx, m)
	})
}

func PersonPassword(next ent.Mutator) ent.Mutator {
	// see: https://datatracker.ietf.org/doc/html/rfc9106#name-parameter-choice
	params := &argon2id.Params{
		Memory:      64 * 1024, // MiB
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16, // 128 bits
		KeyLength:   32, // 256 bits
	}

	return hook.PersonFunc(func(ctx context.Context, m *data.PersonMutation) (ent.Value, error) {
		if pwd, ok := m.Password(); ok {
			hash, err := argon2id.CreateHash(pwd, params)

			if err != nil {
				return nil, err
			}

			m.SetPassword(hash)
		}

		return next.Mutate(ctx, m)
	})
}

func PersonBirthdate(next ent.Mutator) ent.Mutator {
	birthdateValidationError := data.ValidationError{
		Name: "birthdate",
	}

	return hook.PersonFunc(func(ctx context.Context, m *data.PersonMutation) (ent.Value, error) {
		if birthdate, ok := m.Birthdate(); ok {
			if birthdate.After(time.Now()) {
				return nil, fmt.Errorf(
					`data: validator failed for field "Person.birthdate": value cannot be in the future: %w`,
					&birthdateValidationError,
				)
			}
		}

		return next.Mutate(ctx, m)
	})
}

func PersonLanguage(next ent.Mutator) ent.Mutator {
	i18n := language.NewMatcher([]language.Tag{
		language.Catalan, // the first one is for fallback
		language.Spanish,
		language.English,
	})

	return hook.PersonFunc(func(ctx context.Context, m *data.PersonMutation) (ent.Value, error) {
		if lng, ok := m.Language(); ok {
			tag, _ := language.MatchStrings(i18n, lng)

			m.SetLanguage(
				tag.String(),
			)
		}

		return next.Mutate(ctx, m)
	})
}
