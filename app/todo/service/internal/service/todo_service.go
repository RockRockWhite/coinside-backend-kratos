package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/todo"
	"github.com/ljxsteam/coinside-backend-kratos/app/todo/service/internal/data"
	"gorm.io/gorm"
)

type TodoService struct {
	api.UnimplementedTodoServiceServer

	repo data.TodoRepo
}

func (t TodoService) GetTodoById(ctx context.Context, request *api.GetTodoByIdRequest) (*api.GetTodoResponse, error) {
	data, err := t.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetTodoResponse{
			Todo: nil,
			Code: api.Code_ERROR_TODO_NOTFOUND,
		}, nil
	default:
		return &api.GetTodoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	var items []*api.TodoItem
	for _, m := range data.Items {
		items = append(items, &api.TodoItem{
			Content:        m.Content,
			IsFinished:     m.IsFinished,
			FinishedUserId: m.FinishedUserId,
			CreatedAt:      m.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      m.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	todo := &api.TodoInfo{
		Id:        data.Id,
		CardId:    data.CardId,
		Title:     data.Title,
		Items:     items,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetTodoResponse{
		Todo: todo,
		Code: api.Code_OK,
	}, nil
}

func (t TodoService) GetTodoByIdStream(server api.TodoService_GetTodoByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) AddTodo(ctx context.Context, todo *api.TodoInfo) (*api.AddTodoResponse, error) {

	id, err := t.repo.Insert(ctx, &data.Todo{
		Id:     todo.Id,
		CardId: todo.CardId,
		Title:  todo.Title,
	})

	if err != nil {
		return &api.AddTodoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddTodoResponse{Id: id}, nil
}

func (t TodoService) AddTodoStream(server api.TodoService_AddTodoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) SetTodoTitle(ctx context.Context, req *api.SetTodoTitleRequest) (*api.SetTodoTitleResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTodoTitleResponse{
			Code: api.Code_ERROR_TODO_NOTFOUND,
		}, nil
	default:
		return &api.SetTodoTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	one.Title = req.Title

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTodoTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTodoTitleResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TodoService) SetTodoTitleStream(request *api.SetTodoTitleRequest, server api.TodoService_SetTodoTitleStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) DeleteTodo(ctx context.Context, request *api.DeleteTodoRequest) (*api.DeleteTodoResponse, error) {
	if err := t.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteTodoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteTodoResponse{
		Code: api.Code_OK,
	}, nil
}

func (t TodoService) DeleteTodoStream(server api.TodoService_DeleteTodoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) AddItem(ctx context.Context, item *api.TodoItem) (*api.AddItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) AddItemStream(server api.TodoService_AddItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) SetItemContent(ctx context.Context, request *api.SetContentRequest) (*api.SetContentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) SetItemContentStream(server api.TodoService_SetItemContentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) SetItemFinished(ctx context.Context, request *api.SetFinishedRequest) (*api.SetFinishedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) SetItemFinishedStream(server api.TodoService_SetItemFinishedStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) DeleteTodoItem(ctx context.Context, request *api.DeleteTodoItemRequest) (*api.DeleteTodoItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) DeleteTodoItemStream(server api.TodoService_DeleteTodoItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoService) mustEmbedUnimplementedTodoServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewTodoService(repo data.TodoRepo) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
