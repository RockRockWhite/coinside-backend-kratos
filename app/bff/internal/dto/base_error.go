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
)
