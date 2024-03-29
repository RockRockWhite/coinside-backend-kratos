package dto

var (
	ErrorUnauthorized ResponseDto = ResponseDto{
		Code:    "ERROR_UNAUTHORIZED",
		Message: "Unauthorized.",
		Data:    nil,
	}
	ErrorToken ResponseDto = ResponseDto{
		Code:    "ERROR_TOKEN",
		Message: "Token is incorrect or expired.",
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
	ErrorInternal ResponseDto = ResponseDto{
		Code:    "ERROR_INTERNAL",
		Message: "Internal server error.",
		Data:    nil,
	}
)

func NewErrorInternalDto(data interface{}) ResponseDto {
	return ResponseDto{
		Code:    ErrorInternal.Code,
		Message: ErrorInternal.Message,
		Data:    data,
	}
}

func NewOkDto(data interface{}) ResponseDto {
	return ResponseDto{
		Code:    "OK",
		Message: "Success.",
		Data:    data,
	}
}
