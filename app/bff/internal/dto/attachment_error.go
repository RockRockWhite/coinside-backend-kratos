package dto

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/attachment"
)

type AttachmentError map[attachment.Code]ResponseDto

var AttachmentErrorCode AttachmentError

func init() {
	AttachmentErrorCode = AttachmentError{
		attachment.Code_OK: ResponseDto{
			Code:    "OK",
			Message: "Success.",
		},
		attachment.Code_ERROR_UNKNOWN: ResponseDto{
			Code:    "ERROR_UNKNOWN",
			Message: "",
		},
		attachment.Code_ERROR_ATTACHMENT_NOTFOUND: ResponseDto{
			Code:    "ERROR_ATTACHMENT_NOTFOUND",
			Message: "ATTACHMENT not found.",
		},
	}
}
