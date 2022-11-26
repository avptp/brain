package rules

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/avptp/brain/internal/generated/data/privacy"
	"github.com/avptp/brain/internal/transport/request"
)

func FilterPersonRule() privacy.QueryMutationRule {
	type PersonFilter interface {
		WhereID(entql.ValueP)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		viewer := request.ViewerFromCtx(ctx)

		if viewer == nil {
			return privacy.Denyf("unauthenticated")
		}

		filter, ok := f.(PersonFilter)

		if !ok {
			return privacy.Deny
		}

		filter.WhereID(
			entql.ValueEQ(viewer.ID),
		)

		return privacy.Allow
	})
}

func FilterPersonOwnedRule() privacy.QueryMutationRule {
	type PersonOwnedFilter interface {
		WherePersonID(entql.ValueP)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		viewer := request.ViewerFromCtx(ctx)

		if viewer == nil {
			return privacy.Denyf("unauthenticated")
		}

		filter, ok := f.(PersonOwnedFilter)

		if !ok {
			return privacy.Deny
		}

		filter.WherePersonID(
			entql.ValueEQ(viewer.ID),
		)

		return privacy.Allow
	})
}
