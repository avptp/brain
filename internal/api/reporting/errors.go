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

var ErrInternal = &gqlerror.Error{
	Message:    "Internal system error.",
	Extensions: map[string]any{"code": "INTERNAL"},
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
