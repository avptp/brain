package arguments

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func GetFields(ctx context.Context, name string) []string {
	fieldContext := graphql.GetFieldContext(ctx)

	names := []string{}

	for _, arg := range fieldContext.Field.Arguments {
		if arg.Name != name {
			continue
		}

		for _, child := range arg.Value.Children {
			names = append(names, child.Name)
		}
	}

	return names
}
