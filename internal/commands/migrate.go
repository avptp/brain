package commands

import (
	"context"
	"fmt"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data/person"
)

func Migrate(ctn *container.Container) error {
	data := ctn.GetData()

	return data.Debug().Schema.Create(
		context.Background(),
		schema.WithApplyHook(createGenderEnum),
		schema.WithDiffHook(removeGenderDrop),
	)
}

func createGenderEnum(next schema.Applier) schema.Applier {
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
		}, plan.Changes...)

		return next.Apply(ctx, conn, plan)
	})
}

// Atlas tries to remove the gender type, so it becomes necessary
// to remove the drop query before it is run
func removeGenderDrop(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		changes, err := next.Diff(current, desired)

		if err != nil {
			return nil, err
		}

		for i, change := range changes {
			drop, ok := change.(*atlas.DropObject)

			if !ok {
				continue
			}

			enum, ok := drop.O.(*atlas.EnumType)

			if !ok || enum.T != "gender" {
				continue
			}

			changes = append(changes[:i], changes[i+1:]...)
		}

		return changes, nil
	})
}
