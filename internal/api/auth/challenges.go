package auth

import (
	"errors"
	"time"

	"github.com/avptp/brain/internal/api/reporting"
	"github.com/avptp/brain/internal/generated/data"
)

type ChallengeType string

const (
	PasswordChallenge ChallengeType = "password"
	CaptchaChallenge  ChallengeType = "captcha"
)

type ChallengeSpec struct {
	Type   ChallengeType `json:"type"`
	MaxAge time.Duration `json:"maxAge"`
}

func EnforceChallenges(auth *data.Authentication, specs []ChallengeSpec) error {
	now := time.Now()

	if auth == nil {
		return errors.New("auth: challenges: authentication struct must not be nil")
	}

	for _, spec := range specs {
		var last *time.Time

		switch spec.Type {
		case PasswordChallenge:
			last = auth.LastPasswordChallengeAt
		case CaptchaChallenge:
			last = auth.LastCaptchaChallengeAt
		default:
			return errors.New("auth: challenges: unknown challenge type")
		}

		if last == nil || now.Sub(*last) > spec.MaxAge {
			err := reporting.ErrChallenge
			err.Extensions["challenges"] = specs

			return err
		}
	}

	return nil
}
