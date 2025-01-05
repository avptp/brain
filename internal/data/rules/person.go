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
			return DenyUnauthenticated
		}

		filter, ok := f.(PersonFilter)

		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		filter.WhereID(
			entql.ValueEQ(viewer.ID),
		)

		return privacy.Skip
	})
}

func FilterPersonOwnedRule() privacy.QueryMutationRule {
	type PersonOwnedFilter interface {
		WherePersonID(entql.ValueP)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		viewer := request.ViewerFromCtx(ctx)

		if viewer == nil {
			return DenyUnauthenticated
		}

		filter, ok := f.(PersonOwnedFilter)

		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		filter.WherePersonID(
			entql.ValueEQ(viewer.ID),
		)

		return privacy.Skip
	})
}
