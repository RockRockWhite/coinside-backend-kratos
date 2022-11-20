package dto

import "github.com/ljxsteam/coinside-backend-kratos/api/user"

type UserError map[user.Code]ResponseDto

var UserErrorCode UserError

func init() {
	UserErrorCode = UserError{
		user.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		user.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		user.Code_ERROR_USER_NOTFOUND: ResponseDto{
			Code:    "ERROR_USER_NOTFOUND",
			Message: "User not found.",
		},
		user.Code_ERROR_USER_NICKNAME_EXISTS: ResponseDto{
			Code:    "ERROR_USER_NICKNAME_EXISTS",
			Message: "User nickname exists.",
		},
		user.Code_ERROR_USER_MOBILE_EXISTS: ResponseDto{
			Code:    "ERROR_USER_MOBILE_EXISTS",
			Message: "User mobile exists.",
		},
		user.Code_ERROR_USER_PASSWORD: ResponseDto{
			Code:    "ERROR_USER_ID_OR_PASSWORD",
			Message: "Id or password error.",
		},
		user.Code_ERROR_VERIFY_CODE: ResponseDto{
			Code:    "ERROR_VERIFY_CODE",
			Message: "Verify code exists.",
		},
	}
}
