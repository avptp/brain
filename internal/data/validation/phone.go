package validation

import (
	"errors"

	"github.com/nyaruka/phonenumbers"
)

func Phone(s string) error {
	phone, err := phonenumbers.Parse(s, "")

	if err != nil {
		return err
	}

	result := phonenumbers.IsValidNumber(phone)

	if !result {
		return errors.New("value is invalid")
	}

	return nil
}
