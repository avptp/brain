//go:build ignore
// +build ignore

package main

import (
	"github.com/avptp/brain/internal/services/provider"
	"github.com/sarulabs/dingo/v4"
)

func main() {
	err := dingo.GenerateContainerWithCustomPkgName(
		(*provider.Provider)(nil),
		"../internal/generated",
		"container",
	)

	if err != nil {
		panic(err) // unrecoverable situation
	}
}
