package types

import (
	"database/sql/driver"
	"errors"
	"io"

	"github.com/avptp/brain/internal/encoding"
	"github.com/google/uuid"
)

var (
	ErrIDSize              = errors.New("id: bad size when unmarshaling")
	ErrIDInvalidCharacters = errors.New("id: bad characters when unmarshaling")
	ErrIDScanType          = errors.New("id: source value must be a string")
)

type ID uuid.UUID

var ZeroID ID

func ParseID(s string) (ID, error) {
	data, err := encoding.Base32.DecodeString(s)

	if err != nil {
		return ZeroID, ErrIDInvalidCharacters
	}

	if len(data) > 16 {
		return ZeroID, ErrIDSize
	}

	return ID(data), nil
}

func (id ID) String() string {
	return encoding.Base32.EncodeToString(id[:])
}

// Scan implements sql.Scanner
func (id *ID) Scan(src any) error {
	uuid := uuid.UUID{}
	uuid.Scan(src)

	*id = ID(uuid)

	return nil
}

// Value implements sql.Valuer
func (id ID) Value() (driver.Value, error) {
	uuid := uuid.UUID(id)

	return uuid.String(), nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (id *ID) UnmarshalGQL(src any) error {
	switch src := src.(type) {
	case string:
		dst, err := ParseID(src)

		if err != nil {
			return err
		}

		*id = dst
	default:
		return ErrIDScanType
	}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (id ID) MarshalGQL(w io.Writer) {
	src := id[:]
	dst := make([]byte, encoding.Base32.EncodedLen(len(src)))

	encoding.Base32.Encode(dst, src)

	w.Write([]byte(`"`))
	w.Write(dst)
	w.Write([]byte(`"`))
}
