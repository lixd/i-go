package ret

import "net/http"

const (
	Success    = http.StatusOK
	SuccessMsg = "success"

	Fail       = http.StatusBadRequest
	FailMsg    = "fail"
	BadRequest = "bad request"

	Unauthorized    = http.StatusUnauthorized
	UnauthorizedMsg = "unauthorized"

	Forbidden    = http.StatusForbidden
	ForbiddenMsg = "forbidden"

	Limit    = http.StatusTooManyRequests
	LimitMsg = "too many requests"
)
