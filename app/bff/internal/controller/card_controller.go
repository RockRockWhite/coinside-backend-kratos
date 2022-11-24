package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"net/http"
	"strconv"
)

type CardController struct {
	userClient user.UserClient
	teamClient team.TeamClient
	cardClient card.CardClient
}

func (u *CardController) GetCardInfo(c *gin.Context) {
	cardInfo := c.MustGet("card_info").(*card.CardInfo)

	// 获取冗余用户信息
	type MemberInfo struct {
		*card.CardMember
		Nickname string `json:"nickname"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}

	var members []MemberInfo
	// 获取成员信息
	stream, err := u.userClient.GetUserInfoStream(context.Background())
	defer stream.CloseSend()
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	for _, m := range cardInfo.Members {
		if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		userInfo, err := stream.Recv()
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		members = append(members, MemberInfo{
			CardMember: m,
			Nickname:   userInfo.Info.Nickname,
			Fullname:   userInfo.Info.Fullname,
			Email:      userInfo.Info.Email,
			Avatar:     userInfo.Info.Avatar,
		})
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[card.Code_OK].Code,
		Message: dto.CardErrorCode[card.Code_OK].Message,
		Data: struct {
			*card.CardInfo
			Members []MemberInfo `json:"members"`
		}{
			CardInfo: cardInfo,
			Members:  members,
		},
	})
}

func (u *CardController) GetCardInfoList(c *gin.Context) {
	teamId := c.Query("team_id")
	status := c.Query("status")
	memberIds := c.QueryArray("member_id")
	tags := c.QueryArray("tag")

	claims := c.MustGet("claims").(*util.JwtClaims)

	// 生成过滤器参数
	var filters []*card.CardFilter
	if teamId != "" {
		// 判断必须是该team的leader
		id, err := strconv.ParseUint(teamId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		res, err := u.teamClient.GetTeamById(context.Background(), &team.GetTeamByIdRequest{
			Id: id,
		})
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		switch res.Code {
		case team.Code_OK:
			ok := false
			for _, m := range res.Team.Members {
				if m.IsAdmin && m.UserId == claims.Id {
					ok = true
					break
				}
			}
			if !ok {
				c.JSON(http.StatusOK, dto.ErrorForbidden)
				return
			}

		default:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.TeamErrorCode[res.Code].Code,
				Message: dto.TeamErrorCode[res.Code].Message,
				Data:    nil,
			})
			return
		}

		filters = append(filters, &card.CardFilter{
			Type:  card.CardFilterType_TEAM,
			Value: teamId,
		})
	}

	for _, id := range memberIds {
		// 如果未设置team过滤器，则只允许过滤自身
		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}
		if teamId == "" && userId != claims.Id {
			c.JSON(http.StatusOK, dto.ErrorForbidden)
			return
		}
		filters = append(filters, &card.CardFilter{
			Type:  card.CardFilterType_MEMBER,
			Value: id,
		})
	}
	// 判断上述两种过滤器必须有至少有一个
	if len(filters) == 0 {
		c.JSON(http.StatusOK, dto.CardFilterError)
		return
	}

	if status != "" {
		filters = append(filters, &card.CardFilter{
			Type:  card.CardFilterType_STATUS,
			Value: status,
		})
	}
	for _, tag := range tags {
		filters = append(filters, &card.CardFilter{
			Type:  card.CardFilterType_TAG,
			Value: tag,
		})
	}

	limit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)
	if limit == 0 {
		limit = 20
	}
	offset, _ := strconv.ParseUint(c.Query("offset"), 10, 64)

	res, err := u.cardClient.GetCardInfoList(context.Background(), &card.GetCardInfoListRequest{
		Limit:   limit,
		Offset:  offset,
		Filters: filters,
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	// 获取冗余用户信息
	type MemberInfo struct {
		*card.CardMember
		Nickname string `json:"nickname"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}
	var data []struct {
		*card.CardInfo
		Members []MemberInfo `json:"members"`
	}

	infos := res.Infos
	for _, info := range infos {
		var members []MemberInfo
		// 获取成员信息
		stream, err := u.userClient.GetUserInfoStream(context.Background())
		defer stream.CloseSend()
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		for _, m := range info.Members {
			if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			userInfo, err := stream.Recv()
			if err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			members = append(members, MemberInfo{
				CardMember: m,
				Nickname:   userInfo.Info.Nickname,
				Fullname:   userInfo.Info.Fullname,
				Email:      userInfo.Info.Email,
				Avatar:     userInfo.Info.Avatar,
			})
		}

		data = append(data, struct {
			*card.CardInfo
			Members []MemberInfo `json:"members"`
		}{
			CardInfo: info,
			Members:  members,
		})
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[card.Code_OK].Code,
		Message: dto.CardErrorCode[card.Code_OK].Message,
		Data: struct {
			Count uint64 `json:"count,omitempty"`
			Infos interface{}
		}{
			Count: res.Count,
			Infos: data,
		},
	})
}

func (u *CardController) CreateCard(c *gin.Context) {
	var req card.CreateCardRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}
	req.CreatorId = c.MustGet("claims").(*util.JwtClaims).Id

	res, err := u.cardClient.CreateCard(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
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

func (u *CardController) SetTitle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.UpdateCardTitle(context.Background(), &card.UpdateCardTitleRequest{
		Id:    id,
		Title: reqDto.Title,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetContent(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Content string `json:"content"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.UpdateCardContent(context.Background(), &card.UpdateCardContentRequest{
		Id:      id,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetDeadline(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Deadline string `json:"deadline"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.SetCardDeadline(context.Background(), &card.SetCardDeadlineRequest{
		Id:       id,
		Deadline: reqDto.Deadline,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) SetStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Status card.CardStatus `json:"status"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.SetCardStatus(context.Background(), &card.SetCardStatusRequest{
		Id:     id,
		Status: reqDto.Status,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) AddTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagContent := c.Param("tag_content")

	res, err := u.cardClient.AddCardTag(context.Background(), &card.AddCardTagRequest{
		Id:      id,
		Content: tagContent,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) DeleleTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagContent := c.Param("tag_content")

	res, err := u.cardClient.DeleteCardTag(context.Background(), &card.DeleteCardTagRequest{
		Id:      id,
		Content: tagContent,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	reqDto := struct {
		IsAdmin bool `json:"is_admin"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	// 判断用户是否存在
	if res, err := u.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: userId}); err != nil {
		// error
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
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

	// 设置团队成员
	res, err := u.cardClient.SetCardMember(context.Background(), &card.SetCardMemberRequest{
		Id:      id,
		UserId:  userId,
		IsAdmin: reqDto.IsAdmin,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) DeleteMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	// 如果只剩一个Card Member，则拒绝
	cardInfo := c.MustGet("card_info").(*card.CardInfo)
	if len(cardInfo.Members) <= 1 {
		c.JSON(http.StatusOK, &dto.CardOnlyOneMember)
		return
	}

	res, err := u.cardClient.DeleteCardMember(context.Background(), &card.DeleteCardMemberRequest{
		Id:     id,
		UserId: userId,
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) DeleteCard(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.cardClient.DeleteCard(context.Background(), &card.DeleteCardRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) IsCardMember(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		userId := c.MustGet("claims").(*util.JwtClaims).Id

		// 获取当前用户
		cardInfo, err := u.cardClient.GetCardInfo(context.Background(), &card.GetCardInfoRequest{
			Id: id,
		})
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			c.Abort()
			return
		}

		switch cardInfo.Code {
		case card.Code_OK:
			// 判断是否是团队成员
			for _, m := range cardInfo.Info.Members {
				if m.UserId == userId {
					if isAdmin && !m.IsAdmin {
						continue
					}
					// Claims写入上下文
					c.Set("card_info", cardInfo.Info)
					return
				}
			}
			c.JSON(http.StatusOK, &dto.ErrorForbidden)
			c.Abort()
		default:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.CardErrorCode[cardInfo.Code].Code,
				Message: dto.CardErrorCode[cardInfo.Code].Message,
				Data:    nil,
			})
			c.Abort()
		}
	}
}

func NewCardController(userClient user.UserClient, cardClient card.CardClient, teamClient team.TeamClient) *CardController {
	return &CardController{userClient: userClient, cardClient: cardClient, teamClient: teamClient}
}
