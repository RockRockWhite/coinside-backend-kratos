package dto

var (
	ErrorUnauthorized ResponseDto = ResponseDto{
		Code:    "ERROR_UNAUTHORIZED",
		Message: "Unauthorized.",
		Data:    nil,
	}
	ErrorBadRequest ResponseDto = ResponseDto{
		Code:    "ERROR_BAD_REQUEST",
		Message: "Bad request parameters or illegal request.",
		Data:    nil,
	}
	ErrorForbidden ResponseDto = ResponseDto{
		Code:    "ERROR_FORBIDDEN",
		Message: "Cannot not access this resource.",
		Data:    nil,
	}
)
