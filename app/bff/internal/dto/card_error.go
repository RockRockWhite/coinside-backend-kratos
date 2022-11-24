package dto

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
)

type CardError map[card.Code]ResponseDto

var CardErrorCode CardError
var CardOnlyOneMember ResponseDto = ResponseDto{
	Code:    "ERROR_CARD_OLNY_MEMBER",
	Message: "Cannot delete the only member of the card.",
	Data:    nil,
}
var CardFilterError ResponseDto = ResponseDto{
	Code:    "ERROR_CARD_Filter",
	Message: "Must has at lease one query of member_id or team_id.",
	Data:    nil,
}

func init() {
	CardErrorCode = CardError{
		card.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		card.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		card.Code_ERROR_CARD_NOTFOUND: ResponseDto{
			Code:    "ERROR_CARD_NOTFOUND",
			Message: "Card not found.",
		},
	}
}
