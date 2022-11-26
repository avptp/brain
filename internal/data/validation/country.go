package validation

import (
	"errors"

	"github.com/pariz/gountries"
)

func Country(s string) error {
	_, err := gountries.New().FindCountryByAlpha(s)

	if err != nil {
		return errors.New("value is invalid")
	}

	return nil
}
