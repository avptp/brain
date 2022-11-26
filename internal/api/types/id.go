package types

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func MarshalID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		v, _ := ulid.ULID(id).MarshalText()

		w.Write([]byte("\""))
		w.Write(v)
		w.Write([]byte("\""))
	})
}

func UnmarshalID(src any) (uuid.UUID, error) {
	id := ulid.ULID{}
	err := id.Scan(src)

	return uuid.UUID(id), err
}
