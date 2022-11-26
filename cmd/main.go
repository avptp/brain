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
		panic(err) // logger is not yet available
	}

	log := ctn.GetLogger()

	defer func() {
		err := ctn.Delete()

		if err != nil {
			log.Fatal(err)

			// flush buffers again, since container has just been deleted
			// intentionally ignoring error here, see https://github.com/uber-go/zap/issues/328
			_ = log.Sync()
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
		log.Fatal(err)
	}
}
