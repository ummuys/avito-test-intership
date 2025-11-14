package errs

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
	ErrCodeBadToken = "BAD_TOKEN"
	ErrMsgBadToken  = "invalid or expire token"
)
