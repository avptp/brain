package validation

import (
	"errors"

	nif "github.com/criptalia/spanish_dni_validator"
)

func TaxID(v string) error {
	if !nif.IsValid(v) {
		return errors.New("value is invalid")
	}

	return nil
}
