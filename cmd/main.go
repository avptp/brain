package main

import (
	"os"

	"github.com/avptp/brain/internal/commands"
	"github.com/avptp/brain/internal/generated/container"
	_ "github.com/avptp/brain/internal/generated/data/runtime"
)

//go:generate go run -mod=mod ../internal/data/build.go
//go:generate go run -mod=readonly ../internal/services/provider/build.go
//go:generate go run github.com/99designs/gqlgen generate

func main() {
	// Create the dependency container
	ctn, err := container.NewContainer()

	if err != nil {
		panic(err) // unrecoverable situation
	}

	defer func() {
		err := ctn.Delete()

		if err != nil {
			panic(err) // unrecoverable situation
		}
	}()

	// Get command argument
	arg := ""

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	// Run command
	switch arg {
	case "", "serve":
		err = commands.Serve(ctn)
	case "database:migrate":
		err = commands.Migrate(ctn)
	}

	if err != nil {
		panic(err) // unrecoverable situation
	}
}
