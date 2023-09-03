package reporting

import "github.com/vektah/gqlparser/v2/gqlerror"

var ErrCaptcha = &gqlerror.Error{
	Message:    "The captcha challenge has not been met.",
	Extensions: map[string]any{"code": "CAPTCHA"},
}

var ErrConstraint = &gqlerror.Error{
	Message:    "Input violates data constraints.",
	Extensions: map[string]any{"code": "CONSTRAINT"},
}

var ErrInput = &gqlerror.Error{
	Message:    "The input data is not in a valid format.",
	Extensions: map[string]any{"code": "INPUT"},
}

var ErrInternal = &gqlerror.Error{
	Message:    "Internal system error.",
	Extensions: map[string]any{"code": "INTERNAL"},
}

var ErrRateLimit = &gqlerror.Error{
	Message:    "A limit of attempts per unit of time has been reached for this action.",
	Extensions: map[string]any{"code": "RATE_LIMIT"},
}

var ErrNotFound = &gqlerror.Error{
	Message:    "The specified resource does not exist.",
	Extensions: map[string]any{"code": "NOT_FOUND"},
}

var ErrUnauthenticated = &gqlerror.Error{
	Message:    "This action requires authentication.",
	Extensions: map[string]any{"code": "UNAUTHENTICATED"},
}

var ErrUnauthorized = ErrNotFound // not to expose whether a resource exists

var ErrValidation = &gqlerror.Error{
	Message:    "Input does not meet data validation rules.",
	Extensions: map[string]any{"code": "VALIDATION"},
}

var ErrWrongPassword = &gqlerror.Error{
	Message:    "This action requires providing the right password.",
	Extensions: map[string]any{"code": "WRONG_PASSWORD"},
}
