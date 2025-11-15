package errs

import "errors"

// #nosec G101 -- error code, not credentials
var ErrInvalidCredentials = errors.New("invalid credentials")

const (
	ErrCodeUserExists = "USER_EXISTS"
	ErrMsgUserExists  = "user already exists"
)

const (
	ErrCodeNotFound = "NOT_FOUND"
	ErrMsgNotFound  = "resourse not found"
)

const (
	ErrCodeBadJSON = "BAD_JSON"
	ErrMsgBadJSON  = "bad json, please check fields"
)

const (
	ErrCodeInternal = "INTERNAL"
	ErrMsgInternal  = "Something wrong with server, try again later"
)

const (
	// #nosec G101 -- error code, not credentials
	ErrCodeBadToken = "BAD_TOKEN"
	// #nosec G101 -- error code, not credentials
	ErrMsgBadToken = "invalid or expire token"
)

const (
	// #nosec G101 -- error code, not credentials
	ErrCodeInvalidCredentials = "INVALID_CREDENTIALS"
	// #nosec G101 -- error code, not credentials
	ErrMsgInvalidCredentials = "username or password incorrect"
)
