//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithRelaySpec(true),
	)

	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.FeatureNames(
			gen.FeaturePrivacy.Name,
			gen.FeatureEntQL.Name,
			gen.FeatureExecQuery.Name,
		),
		entc.Extensions(ex),
		entc.TemplateDir("../internal/data/templates"),
	}

	err = entc.Generate("../internal/data/schema", &gen.Config{
		Target:    "../internal/generated/data",
		Package:   "github.com/avptp/brain/internal/generated/data",
		Templates: entgql.AllTemplates,
	}, opts...)

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
