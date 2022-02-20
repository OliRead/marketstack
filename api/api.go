package api

import "errors"

var (
	ErrTransport                = errors.New("transport error")
	ErrDecode                   = errors.New("decode error")
	ErrInternal                 = errors.New("internal error")
	ErrInvalidAccessKey         = errors.New("invalid access key")
	ErrMissingAccessKey         = errors.New("missing access key")
	ErrInactiveUse              = errors.New("inactive user")
	ErrHTTPSAccessRestricted    = errors.New("https access restricted")
	ErrFunctionAccessRestricted = errors.New("function access restricted")
	ErrInvalidAPIFunction       = errors.New("invalid API function")
	ErrNotFound                 = errors.New("404 not found")
	ErrUsageLimitReached        = errors.New("usage limit reached")
	ErrRateLimitReached         = errors.New("rate limit reached")
)
