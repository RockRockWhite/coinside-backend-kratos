package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/api/vote"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type VoteController struct {
	userClient user.UserClient
	voteClient vote.VoteClient
}

func (t *VoteController) GetVoteInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.voteClient.GetVoteById(context.Background(), &vote.GetVoteByIdRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	switch res.Code {
	case vote.Code_OK:
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
		//for _, m := range res.Vote.Members {
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
		//		VoteMember: m,
		//		Nickname:   userInfo.Info.Nickname,
		//		Fullname:   userInfo.Info.Fullname,
		//		Email:      userInfo.Info.Email,
		//		Avatar:     userInfo.Info.Avatar,
		//	})
		//}

		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.VoteErrorCode[res.Code].Code,
			Message: dto.VoteErrorCode[res.Code].Message,
			Data: struct {
				*vote.VoteInfo
				//Members []MemberInfo `json:"members"`
			}{
				VoteInfo: res.Vote,
				//Members: members,
			},
		})

	default:
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.VoteErrorCode[res.Code].Code,
			Message: dto.VoteErrorCode[res.Code].Message,
			Data:    nil,
		})
	}
	//
	//if res.Code != team.Code_OK {
	//	resDto.Data = err
	//} else {
	//	resDto.Data = res.Vote
	//}
	//
	//c.JSON(http.StatusOK, resDto)
}

func (t *VoteController) CreateVote(c *gin.Context) {
	var req vote.AddVoteRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}
	//req.CreatorId = c.MustGet("claims").(*util.JwtClaims).Id

	res, err := t.voteClient.AddVote(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
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

func (t *VoteController) SetTitle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.voteClient.SetVoteTitle(context.Background(), &vote.SetVoteTitleRequest{
		Id:    id,
		Title: reqDto.Title,
	})

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *VoteController) DeleteVote(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.voteClient.DeleteVote(context.Background(), &vote.DeleteVoteRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *VoteController) SetVoteItem(c *gin.Context) {
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

	res, err := t.voteClient.AddItem(context.Background(), &vote.AddItemRequest{
		VoteId:  id,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
		resDto.Data = err
	} else {

		resDto.Data = struct {
			ItemId uint64 `json:"item_id,omitempty"`
		}{
			ItemId: res.Id,
		}
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *VoteController) SetItemContent(c *gin.Context) {
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

	res, err := t.voteClient.SetItemContent(context.Background(), &vote.SetContentRequest{
		Id:      id,
		ItemId:  itemId,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *VoteController) DeleteVoteItem(c *gin.Context) {
	voteId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	voteItemId, _ := strconv.ParseUint(c.Param("item_id"), 10, 64)

	res, err := t.voteClient.DeleteVoteItem(context.Background(), &vote.DeleteVoteItemRequest{
		VoteId:     voteId,
		VoteItemId: voteItemId,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	resDto := dto.ResponseDto{
		Code:    dto.VoteErrorCode[res.Code].Code,
		Message: dto.VoteErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != vote.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewVoteController(userClient user.UserClient, client vote.VoteClient) *VoteController {
	return &VoteController{userClient: userClient, voteClient: client}
}
