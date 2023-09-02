package commands

import (
	"context"
	"fmt"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data/authorization"
	"github.com/avptp/brain/internal/generated/data/person"
)

func Migrate(ctx context.Context, ctn *container.Container) error {
	data := ctn.GetData()

	return data.Debug().Schema.Create(
		ctx,
		schema.WithApplyHook(createEnums),
		schema.WithDiffHook(keepEnums),
	)
}

func createEnums(next schema.Applier) schema.Applier {
	return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
		plan.Changes = append([]*migrate.Change{
			{
				Cmd: fmt.Sprintf(
					"CREATE TYPE IF NOT EXISTS gender AS ENUM ('%s', '%s', '%s');",
					person.GenderWoman,
					person.GenderMan,
					person.GenderNonBinary,
				),
			},
			{
				Cmd: fmt.Sprintf(
					"CREATE TYPE IF NOT EXISTS authorization_kind AS ENUM ('%s', '%s');",
					authorization.KindEmail,
					authorization.KindPassword,
				),
			},
		}, plan.Changes...)

		return next.Apply(ctx, conn, plan)
	})
}

// Atlas tries to remove the enum types, so it becomes necessary
// to add them to the desired objects slice
func keepEnums(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		for _, object := range current.Objects {
			_, ok := object.(*atlas.EnumType)

			if !ok {
				continue
			}

			desired.Objects = append(desired.Objects, object)
		}

		return next.Diff(current, desired)
	})
}
