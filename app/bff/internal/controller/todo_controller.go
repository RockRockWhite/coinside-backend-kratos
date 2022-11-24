package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/todo"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type TodoController struct {
	userClient user.UserClient
	todoClient todo.TodoServiceClient
}

func (t *TodoController) GetTodoInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.todoClient.GetTodoById(context.Background(), &todo.GetTodoByIdRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	switch res.Code {
	case todo.Code_OK:
		//// 获取冗余用户信息
		//type MemberInfo struct {
		//	*team.TeamMember
		//	Nickname string `json:"nickname"`
		//	Fullname string `json:"fullname"`
		//	Email    string `json:"email"`
		//	Avatar   string `json:"avatar"`
		//}
		//
		//var members []
		//// 获取成员信息
		//stream, err := t.userClient.GetUserInfoStream(context.Background())
		//defer stream.CloseSend()
		//if err != nil {
		//	c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		//	return
		//}
		//
		//for _, m := range res.Todo.Members {
		//	if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
		//		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		//		return
		//	}
		//
		//	userInfo, err := stream.Recv()
		//	if err != nil {
		//		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		//		return
		//	}
		//
		//	members = append(members, MemberInfo{
		//		TodoMember: m,
		//		Nickname:   userInfo.Info.Nickname,
		//		Fullname:   userInfo.Info.Fullname,
		//		Email:      userInfo.Info.Email,
		//		Avatar:     userInfo.Info.Avatar,
		//	})
		//}

		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TodoErrorCode[res.Code].Code,
			Message: dto.TodoErrorCode[res.Code].Message,
			Data: struct {
				*todo.TodoInfo
				//Members []MemberInfo `json:"members"`
			}{
				TodoInfo: res.Todo,
				//Members: members,
			},
		})

	default:
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TodoErrorCode[res.Code].Code,
			Message: dto.TodoErrorCode[res.Code].Message,
			Data:    nil,
		})
	}
	//
	//if res.Code != team.Code_OK {
	//	resDto.Data = err
	//} else {
	//	resDto.Data = res.Todo
	//}
	//
	//c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) CreateTodo(c *gin.Context) {
	var req todo.AddTodoRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}
	//req.CreatorId = c.MustGet("claims").(*util.JwtClaims).Id

	res, err := t.todoClient.AddTodo(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = struct {
			Id uint64 `json:"id"`
		}{
			Id: res.Id,
		}
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) SetTitle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.todoClient.SetTodoTitle(context.Background(), &todo.SetTodoTitleRequest{
		Id:    id,
		Title: reqDto.Title,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) DeleteTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.todoClient.DeleteTodo(context.Background(), &todo.DeleteTodoRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) SetTodoItem(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	reqDto := struct {
		Content string `json:"content"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	//// 判断用户是否存在
	//if res, err := t.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: userId}); err != nil {
	//	// error
	//	c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
	//	return
	//} else {
	//	// no error
	//	switch res.Code {
	//	case user.Code_OK:
	//	case user.Code_ERROR_USER_NOTFOUND:
	//		c.JSON(http.StatusOK, &dto.ResponseDto{
	//			Code:    dto.UserErrorCode[res.Code].Code,
	//			Message: dto.UserErrorCode[res.Code].Message,
	//			Data:    nil,
	//		})
	//		return
	//	default:
	//		c.JSON(http.StatusOK, &dto.ResponseDto{
	//			Code:    dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Code,
	//			Message: dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Message,
	//			Data:    err,
	//		})
	//		return
	//	}
	//}

	res, err := t.todoClient.AddItem(context.Background(), &todo.TodoItem{
		TodoId:  id,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) SetItemContent(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	itemId, _ := strconv.ParseUint(c.Param("item_id"), 10, 64)
	//userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)
	//
	reqDto := struct {
		Content string `json:"content"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}
	//
	//// 判断用户是否存在
	//if res, err := t.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: userId}); err != nil {
	//	// error
	//	c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
	//	return
	//} else {
	//	// no error
	//	switch res.Code {
	//	case user.Code_OK:
	//	case user.Code_ERROR_USER_NOTFOUND:
	//		c.JSON(http.StatusOK, &dto.ResponseDto{
	//			Code:    dto.UserErrorCode[res.Code].Code,
	//			Message: dto.UserErrorCode[res.Code].Message,
	//			Data:    nil,
	//		})
	//		return
	//	default:
	//		c.JSON(http.StatusOK, &dto.ResponseDto{
	//			Code:    dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Code,
	//			Message: dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Message,
	//			Data:    err,
	//		})
	//		return
	//	}
	//}

	res, err := t.todoClient.SetItemContent(context.Background(), &todo.SetContentRequest{
		Id:      id,
		ItemId:  itemId,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) SetItemFinished(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	itemId, _ := strconv.ParseUint(c.Param("item_id"), 10, 64)
	//userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	reqDto := struct {
		IsFinished bool   `json:"is_finished"`
		UserId     uint64 `json:"finished_user_id"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	// 判断用户是否存在
	if res, err := t.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: reqDto.UserId}); err != nil {
		// error
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	} else {
		// no error
		switch res.Code {
		case user.Code_OK:
		case user.Code_ERROR_USER_NOTFOUND:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.UserErrorCode[res.Code].Code,
				Message: dto.UserErrorCode[res.Code].Message,
				Data:    nil,
			})
			return
		default:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Code,
				Message: dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Message,
				Data:    err,
			})
			return
		}
	}

	res, err := t.todoClient.SetItemFinished(context.Background(), &todo.SetFinishedRequest{
		Id:         id,
		ItemId:     itemId,
		IsFinished: reqDto.IsFinished,
		UserId:     reqDto.UserId,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TodoController) DeleteTodoItem(c *gin.Context) {
	todoId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	todoItemId, _ := strconv.ParseUint(c.Param("item_id"), 10, 64)

	res, err := t.todoClient.DeleteTodoItem(context.Background(), &todo.DeleteTodoItemRequest{
		TodoId:     todoId,
		TodoItemId: todoItemId,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	resDto := dto.ResponseDto{
		Code:    dto.TodoErrorCode[res.Code].Code,
		Message: dto.TodoErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != todo.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewTodoController(userClient user.UserClient, client todo.TodoServiceClient) *TodoController {
	return &TodoController{userClient: userClient, todoClient: client}
}
