package reporting

import "github.com/vektah/gqlparser/v2/gqlerror"

var ErrCaptcha = &gqlerror.Error{
	Message:    "The captcha challenge has not been met.",
	Extensions: map[string]any{"code": "CAPTCHA"},
}

var ErrInternal = &gqlerror.Error{
	Message:    "Internal system error.",
	Extensions: map[string]any{"code": "INTERNAL"},
}

var ErrUnauthenticated = &gqlerror.Error{
	Message:    "This action requires authentication.",
	Extensions: map[string]any{"code": "UNAUTHENTICATED"},
}

var ErrWrongPassword = &gqlerror.Error{
	Message:    "This action requires providing the right password.",
	Extensions: map[string]any{"code": "WRONG_PASSWORD"},
}
