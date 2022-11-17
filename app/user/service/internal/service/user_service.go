package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/data"
	utils "github.com/ljxsteam/coinside-backend-kratos/pkg/util"
	"gorm.io/gorm"
)

type UserService struct {
	api.UnimplementedUserServer

	repo data.UserRepo
}

func (u UserService) CreateUser(ctx context.Context, request *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	salt := utils.GenerateSalt()
	passwordHash := utils.EncryptPasswordHash(request.Passwd, salt)

	if _, err := u.repo.FindOneByNickname(ctx, request.Nickname); err != gorm.ErrRecordNotFound {
		return &api.CreateUserResponse{
			Code: api.Code_ERROR_USER_NICKNAME_EXISTS,
		}, nil
	}

	if _, err := u.repo.FindOneByEmail(ctx, request.Email); err != gorm.ErrRecordNotFound {
		return &api.CreateUserResponse{
			Code: api.Code_ERROR_USER_EMAIL_EXISTS,
		}, nil
	}

	if _, err := u.repo.FindOneByMobile(ctx, request.Mobile); err != gorm.ErrRecordNotFound {
		return &api.CreateUserResponse{
			Code: api.Code_ERROR_USER_MOBILE_EXISTS,
		}, nil
	}

	id, err := u.repo.Insert(ctx, &data.User{
		Nickname:   request.Nickname,
		PasswdSalt: salt,
		PasswdHash: passwordHash,
		Email:      request.Email,
		Mobile:     request.Mobile,
	})

	if err != nil {
		return &api.CreateUserResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.CreateUserResponse{Id: id}, nil
}

func (u UserService) CreateUserStream(server api.User_CreateUserStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetUserInfoResponse{
			Info: nil,
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.GetUserInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &api.UserInfo{
		Id:            one.Id,
		Nickname:      one.Nickname,
		Fullname:      one.Fullname,
		Avatar:        one.Avatar,
		Email:         one.Email,
		EmailVerified: one.EmailVerified,
		Mobile:        one.Mobile,
		Config:        one.Config,
		LoginedAt:     one.LoginedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:     one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetUserInfoResponse{
		Info: info,
		Code: api.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoStream(server api.User_GetUserInfoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByNickname(ctx context.Context, request *api.GetUserInfoByNicknameRequest) (*api.GetUserInfoResponse, error) {
	one, err := u.repo.FindOneByNickname(ctx, request.Nickname)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetUserInfoResponse{
			Info: nil,
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.GetUserInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &api.UserInfo{
		Id:            one.Id,
		Nickname:      one.Nickname,
		Fullname:      one.Fullname,
		Avatar:        one.Avatar,
		Email:         one.Email,
		EmailVerified: one.EmailVerified,
		Mobile:        one.Mobile,
		Config:        one.Config,
		LoginedAt:     one.LoginedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:     one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetUserInfoResponse{
		Info: info,
		Code: api.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByNicknameStream(server api.User_GetUserInfoByNicknameStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByEmail(ctx context.Context, request *api.GetUserInfoByEmailRequest) (*api.GetUserInfoResponse, error) {

	one, err := u.repo.FindOneByEmail(ctx, request.Email)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetUserInfoResponse{
			Info: nil,
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.GetUserInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &api.UserInfo{
		Id:            one.Id,
		Nickname:      one.Nickname,
		Fullname:      one.Fullname,
		Avatar:        one.Avatar,
		Email:         one.Email,
		EmailVerified: one.EmailVerified,
		Mobile:        one.Mobile,
		Config:        one.Config,
		LoginedAt:     one.LoginedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:     one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetUserInfoResponse{
		Info: info,
		Code: api.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByEmailStream(server api.User_GetUserInfoByEmailStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByMobile(ctx context.Context, request *api.GetUserInfoByMobileRequest) (*api.GetUserInfoResponse, error) {
	one, err := u.repo.FindOneByMobile(ctx, request.Mobile)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetUserInfoResponse{
			Info: nil,
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.GetUserInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &api.UserInfo{
		Id:            one.Id,
		Nickname:      one.Nickname,
		Fullname:      one.Fullname,
		Avatar:        one.Avatar,
		Email:         one.Email,
		EmailVerified: one.EmailVerified,
		Mobile:        one.Mobile,
		Config:        one.Config,
		LoginedAt:     one.LoginedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:     one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetUserInfoResponse{
		Info: info,
		Code: api.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByMobileStream(server api.User_GetUserInfoByMobileStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetFullname(ctx context.Context, request *api.SetFullnameRequest) (*api.SetFullnameResponse, error) {

	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetFullnameResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.SetFullnameResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Fullname = request.Fullname
	if err = u.repo.Update(ctx, one); err != nil {
		return &api.SetFullnameResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.SetFullnameResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) SetFullnameStream(server api.User_SetFullnameStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetAvatar(ctx context.Context, request *api.SetAvatarRequest) (*api.SetAvatarResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetAvatarResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.SetAvatarResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Avatar = request.Avatar
	if err = u.repo.Update(ctx, one); err != nil {
		return &api.SetAvatarResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.SetAvatarResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) SetAvatarStream(server api.User_SetAvatarStreamServer) error {

	//TODO implement me
	panic("implement me")
}

func (u UserService) SetConfig(ctx context.Context, request *api.SetConfigRequest) (*api.SetConfigResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetConfigResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.SetConfigResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Config = request.Config
	if err = u.repo.Update(ctx, one); err != nil {
		return &api.SetConfigResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.SetConfigResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) SetConfigStream(server api.User_SetConfigStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetEmail(ctx context.Context, request *api.SetEmailRequest) (*api.SetEmailResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetEmailResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.SetEmailResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	// todo: check verify_code
	verifyCode := "123456"
	if request.VerifyCode != verifyCode {
		return &api.SetEmailResponse{
			Code: api.Code_ERROR_VERIFY_CODE,
		}, nil
	}

	// check email is exists
	if _, err = u.repo.FindOneByEmail(ctx, request.Email); err != gorm.ErrRecordNotFound {
		return &api.SetEmailResponse{
			Code: api.Code_ERROR_USER_EMAIL_EXISTS,
		}, nil
	}

	one.Email = request.Email
	if err = u.repo.Update(ctx, one); err != nil {
		return &api.SetEmailResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.SetEmailResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) SetEmailStream(server api.User_SetEmailStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetMobile(ctx context.Context, request *api.SetMobileRequest) (*api.SetMobileResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetMobileResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.SetMobileResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	// todo: check verify_code
	verifyCode := "123456"
	if request.VerifyCode != verifyCode {
		return &api.SetMobileResponse{
			Code: api.Code_ERROR_VERIFY_CODE,
		}, nil
	}

	// check mobile is exists
	if _, err = u.repo.FindOneByMobile(ctx, request.Mobile); err != gorm.ErrRecordNotFound {
		return &api.SetMobileResponse{
			Code: api.Code_ERROR_USER_MOBILE_EXISTS,
		}, nil
	}

	one.Mobile = request.Mobile
	if err = u.repo.Update(ctx, one); err != nil {
		return &api.SetMobileResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.SetMobileResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) SetMobileStream(server api.User_SetMobileStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteUser(ctx context.Context, request *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	if err := u.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteUserResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteUserResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) DeleteUserStream(server api.User_DeleteUserStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.LoginResponse{
			Code: api.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &api.LoginResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	// check password
	if !utils.ValifyPasswordHash(request.Password, one.PasswdSalt, one.PasswdHash) {
		return &api.LoginResponse{
			Code: api.Code_ERROR_USER_PASSWORD,
		}, nil
	}

	return &api.LoginResponse{
		Code: api.Code_OK,
	}, nil
}

func (u UserService) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repo data.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
