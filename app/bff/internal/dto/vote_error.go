package dto

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/vote"
)

type VoteError map[vote.Code]ResponseDto

var VoteErrorCode VoteError

func init() {
	VoteErrorCode = VoteError{
		vote.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		vote.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		vote.Code_ERROR_VOTE_NOTFOUND: ResponseDto{
			Code:    "ERROR_VOTE_NOTFOUND",
			Message: "VOTE not found.",
		},
	}
}
