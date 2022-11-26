package validation

import (
	"errors"

	nif "github.com/criptalia/spanish_dni_validator"
)

func TaxID(s string) error {
	if !nif.IsValid(s) {
		return errors.New("value is invalid")
	}

	return nil
}
