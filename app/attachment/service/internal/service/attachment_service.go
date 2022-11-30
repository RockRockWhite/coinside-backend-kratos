package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/attachment"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/internal/data"
	"gorm.io/gorm"
)

type AttachmentService struct {
	api.UnimplementedAttachmentServer

	repo data.AttachmentRepo
}

func (a AttachmentService) GetAttachmentById(ctx context.Context, request *api.GetAttachmentByIdRequest) (*api.GetAttachmentResponse, error) {
	data, err := a.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetAttachmentResponse{
			Attachment: nil,
			Code:       api.Code_ERROR_ATTACHMENT_NOTFOUND,
		}, nil
	default:
		return &api.GetAttachmentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	attachment := &api.AttachmentInfo{
		Id:            data.Id,
		CardId:        data.CardId,
		Link:          data.Link,
		DownloadCount: data.DownloadCount,
		CreatedAt:     data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetAttachmentResponse{
		Attachment: attachment,
		Code:       api.Code_OK,
	}, nil
}

func (a AttachmentService) GetAttachmentByIdStream(server api.Attachment_GetAttachmentByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (a AttachmentService) GetAttachmentsByCardId(ctx context.Context, request *api.GetAttachmentsByCardIdRequest) (*api.GetAttachmentsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AttachmentService) GetAttachmentsByCardIdStream(server api.Attachment_GetAttachmentsByCardIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (a AttachmentService) AddAttachment(ctx context.Context, attachment *api.AddAttachmentRequest) (*api.AddAttachmentResponse, error) {
	id, err := a.repo.Insert(ctx, &data.Attachment{
		CardId:        attachment.CardId,
		Link:          attachment.Link,
		DownloadCount: 0,
	})

	if err != nil {
		return &api.AddAttachmentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddAttachmentResponse{Id: id}, nil
}

func (a AttachmentService) AddAttachmentStream(server api.Attachment_AddAttachmentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

// need the method: UpdateAttachment(context.Context, *UpdateAttachmentRequest) (*UpdateAttachmentResponse, error)
func (a AttachmentService) UpdateAttachment(ctx context.Context, attachment *api.UpdateAttachmentRequest) (*api.UpdateAttachmentResponse, error) {
	one, err := a.repo.FindOne(ctx, attachment.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateAttachmentResponse{
			Code: api.Code_ERROR_ATTACHMENT_NOTFOUND,
		}, nil
	default:
		return &api.UpdateAttachmentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	NewAttachment := data.Attachment{
		Id:            one.Id,
		CardId:        one.CardId,
		Link:          one.Link,
		DownloadCount: one.DownloadCount,
		CreatedAt:     one.CreatedAt,
	}

	if attachment.Link != "" && attachment.Link != NewAttachment.Link {
		NewAttachment.Link = attachment.Link
	}

	if attachment.DownloadCount != 0 && attachment.DownloadCount != NewAttachment.DownloadCount {
		NewAttachment.DownloadCount = attachment.DownloadCount
	}

	if error := a.repo.Update(ctx, &NewAttachment); error != nil {
		return &api.UpdateAttachmentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.UpdateAttachmentResponse{
		Code: api.Code_OK,
	}, nil

}

func (a AttachmentService) UpdateAttachmentStream(info *api.UpdateAttachmentRequest, server api.Attachment_UpdateAttachmentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (a AttachmentService) DeleteAttachment(ctx context.Context, request *api.DeleteAttachmentRequest) (*api.DeleteAttachmentResponse, error) {
	if err := a.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteAttachmentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteAttachmentResponse{
		Code: api.Code_OK,
	}, nil
}

func (a AttachmentService) DeleteAttachmentStream(server api.Attachment_DeleteAttachmentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (a AttachmentService) mustEmbedUnimplementedAttachmentServer() {
	//TODO implement me
	panic("implement me")
}

func NewAttachmentService(repo data.AttachmentRepo) *AttachmentService {
	return &AttachmentService{
		repo: repo,
	}
}
