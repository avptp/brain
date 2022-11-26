package commands

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/avptp/brain/internal/generated/container"
	"github.com/avptp/brain/internal/generated/data/migrate"
	"github.com/avptp/brain/internal/generated/data/person"
)

func Migrate(ctn *container.Container) error {
	data := ctn.GetData()

	return data.Debug().Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				// Create enums
				_, err := data.ExecContext(
					ctx,
					fmt.Sprintf(
						"CREATE TYPE gender AS ENUM ('%s', '%s', '%s');",
						person.GenderWoman,
						person.GenderMan,
						person.GenderNonBinary,
					),
				)

				if err != nil {
					return err
				}

				return next.Create(ctx, tables...)
			})
		}),
	)
}
