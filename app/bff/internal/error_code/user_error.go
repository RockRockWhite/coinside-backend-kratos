package error_code

import api "github.com/ljxsteam/coinside-backend-kratos/api/user"

type UserError map[api.Code]Error

var userError UserError

func init() {
	userError = UserError{
		api.Code_OK: Error{
			Code:    "OK",
			Message: "Success.",
		},
		api.Code_ERROR_UNKNOWN: Error{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		api.Code_ERROR_USER_NOTFOUND: Error{
			Code:    "Code_ERROR_USER_NOTFOUND",
			Message: "User not found.",
		},
		api.Code_ERROR_USER_NICKNAME_EXISTS: Error{
			Code:    "ERROR_USER_NICKNAME_EXISTS",
			Message: "User nickname exists.",
		},
		api.Code_ERROR_USER_MOBILE_EXISTS: Error{
			Code:    "ERROR_USER_MOBILE_EXISTS",
			Message: "User mobile exists.",
		},
		api.Code_ERROR_USER_PASSWORD: Error{
			Code:    "ERROR_USER_PASSWORD",
			Message: "Username or password error.",
		},
		api.Code_ERROR_VERIFY_CODE: Error{
			Code:    "ERROR_VERIFY_CODE",
			Message: "Verify code exists.",
		},
	}
}

func GetUserError() UserError {
	return userError
}
