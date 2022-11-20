package dto

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/team"
)

type TeamError map[team.Code]ResponseDto

var TeamErrorCode TeamError

func init() {
	TeamErrorCode = TeamError{
		team.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		team.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		team.Code_ERROR_TEAM_NOTFOUND: ResponseDto{
			Code:    "ERROR_TEAM_NOTFOUND",
			Message: "Card not found.",
		},
	}
}
