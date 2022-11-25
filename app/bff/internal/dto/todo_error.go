package dto

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/todo"
)

type TodoError map[todo.Code]ResponseDto

var TodoErrorCode TodoError

func init() {
	TodoErrorCode = TodoError{
		todo.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		todo.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		todo.Code_ERROR_TODO_NOTFOUND: ResponseDto{
			Code:    "ERROR_TODO_NOTFOUND",
			Message: "TODO not found.",
		},
	}
}
