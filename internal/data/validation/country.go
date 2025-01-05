package validation

import (
	"errors"

	"github.com/pariz/gountries"
)

var Countries = gountries.New()

func Country(v string) error {
	_, err := Countries.FindCountryByAlpha(v)

	if err != nil {
		return errors.New("value is invalid")
	}

	return nil
}
