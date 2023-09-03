package validation

import (
	"errors"

	"github.com/pariz/gountries"
)

func Country(v string) error {
	_, err := gountries.New().FindCountryByAlpha(v)

	if err != nil {
		return errors.New("value is invalid")
	}

	return nil
}
