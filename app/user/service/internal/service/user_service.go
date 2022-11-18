package service

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/data"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/util"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	user.UnimplementedUserServer

	repo data.UserRepo
}

func (u UserService) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	salt := util.GenerateSalt()
	passwordHash := util.EncryptPasswordHash(request.Password, salt)

	if _, err := u.repo.FindOneByNickname(ctx, request.Nickname); err != gorm.ErrRecordNotFound {
		return &user.CreateUserResponse{
			Code: user.Code_ERROR_USER_NICKNAME_EXISTS,
		}, nil
	}

	if _, err := u.repo.FindOneByEmail(ctx, request.Email); err != gorm.ErrRecordNotFound {
		return &user.CreateUserResponse{
			Code: user.Code_ERROR_USER_EMAIL_EXISTS,
		}, nil
	}

	if _, err := u.repo.FindOneByMobile(ctx, request.Mobile); err != gorm.ErrRecordNotFound {
		return &user.CreateUserResponse{
			Code: user.Code_ERROR_USER_MOBILE_EXISTS,
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
		return &user.CreateUserResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	return &user.CreateUserResponse{Id: id}, nil
}

func (u UserService) CreateUserStream(server user.User_CreateUserStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.GetUserInfoResponse{
			Info: nil,
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.GetUserInfoResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &user.UserInfo{
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

	return &user.GetUserInfoResponse{
		Info: info,
		Code: user.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoStream(server user.User_GetUserInfoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByNickname(ctx context.Context, request *user.GetUserInfoByNicknameRequest) (*user.GetUserInfoResponse, error) {
	one, err := u.repo.FindOneByNickname(ctx, request.Nickname)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.GetUserInfoResponse{
			Info: nil,
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.GetUserInfoResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &user.UserInfo{
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

	return &user.GetUserInfoResponse{
		Info: info,
		Code: user.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByNicknameStream(server user.User_GetUserInfoByNicknameStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByEmail(ctx context.Context, request *user.GetUserInfoByEmailRequest) (*user.GetUserInfoResponse, error) {

	one, err := u.repo.FindOneByEmail(ctx, request.Email)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.GetUserInfoResponse{
			Info: nil,
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.GetUserInfoResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &user.UserInfo{
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

	return &user.GetUserInfoResponse{
		Info: info,
		Code: user.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByEmailStream(server user.User_GetUserInfoByEmailStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserInfoByMobile(ctx context.Context, request *user.GetUserInfoByMobileRequest) (*user.GetUserInfoResponse, error) {
	one, err := u.repo.FindOneByMobile(ctx, request.Mobile)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.GetUserInfoResponse{
			Info: nil,
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.GetUserInfoResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &user.UserInfo{
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

	return &user.GetUserInfoResponse{
		Info: info,
		Code: user.Code_OK,
	}, nil
}

func (u UserService) GetUserInfoByMobileStream(server user.User_GetUserInfoByMobileStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetFullname(ctx context.Context, request *user.SetFullnameRequest) (*user.SetFullnameResponse, error) {

	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.SetFullnameResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.SetFullnameResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Fullname = request.Fullname
	if err = u.repo.Update(ctx, one); err != nil {
		return &user.SetFullnameResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	return &user.SetFullnameResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) SetFullnameStream(server user.User_SetFullnameStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetAvatar(ctx context.Context, request *user.SetAvatarRequest) (*user.SetAvatarResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.SetAvatarResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.SetAvatarResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Avatar = request.Avatar
	if err = u.repo.Update(ctx, one); err != nil {
		return &user.SetAvatarResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}
	return &user.SetAvatarResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) SetAvatarStream(server user.User_SetAvatarStreamServer) error {

	//TODO implement me
	panic("implement me")
}

func (u UserService) SetConfig(ctx context.Context, request *user.SetConfigRequest) (*user.SetConfigResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.SetConfigResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.SetConfigResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Config = request.Config
	if err = u.repo.Update(ctx, one); err != nil {
		return &user.SetConfigResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}
	return &user.SetConfigResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) SetConfigStream(server user.User_SetConfigStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetEmail(ctx context.Context, request *user.SetEmailRequest) (*user.SetEmailResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.SetEmailResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.SetEmailResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	// todo: check verify_code
	verifyCode := "123456"
	if request.VerifyCode != verifyCode {
		return &user.SetEmailResponse{
			Code: user.Code_ERROR_VERIFY_CODE,
		}, nil
	}

	// check email is exists
	if _, err = u.repo.FindOneByEmail(ctx, request.Email); err != gorm.ErrRecordNotFound {
		return &user.SetEmailResponse{
			Code: user.Code_ERROR_USER_EMAIL_EXISTS,
		}, nil
	}

	one.Email = request.Email
	if err = u.repo.Update(ctx, one); err != nil {
		return &user.SetEmailResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}
	return &user.SetEmailResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) SetEmailStream(server user.User_SetEmailStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) SetMobile(ctx context.Context, request *user.SetMobileRequest) (*user.SetMobileResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.SetMobileResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.SetMobileResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	// todo: check verify_code
	verifyCode := "123456"
	if request.VerifyCode != verifyCode {
		return &user.SetMobileResponse{
			Code: user.Code_ERROR_VERIFY_CODE,
		}, nil
	}

	// check mobile is exists
	if _, err = u.repo.FindOneByMobile(ctx, request.Mobile); err != gorm.ErrRecordNotFound {
		return &user.SetMobileResponse{
			Code: user.Code_ERROR_USER_MOBILE_EXISTS,
		}, nil
	}

	one.Mobile = request.Mobile
	if err = u.repo.Update(ctx, one); err != nil {
		return &user.SetMobileResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}
	return &user.SetMobileResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) SetMobileStream(server user.User_SetMobileStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	if err := u.repo.Delete(ctx, request.Id); err != nil {
		return &user.DeleteUserResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}
	return &user.DeleteUserResponse{
		Code: user.Code_OK,
	}, nil
}

func (u UserService) DeleteUserStream(server user.User_DeleteUserStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u UserService) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	one, err := u.repo.FindOne(ctx, request.Id)
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &user.LoginResponse{
			Code: user.Code_ERROR_USER_NOTFOUND,
		}, nil

	default:
		return &user.LoginResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	// check password
	if !util.ValifyPasswordHash(request.Password, one.PasswdSalt, one.PasswdHash) {
		return &user.LoginResponse{
			Code: user.Code_ERROR_USER_PASSWORD,
		}, nil
	}

	// login succeed.
	one.LoginedAt = time.Now()
	err = u.repo.Update(ctx, one)
	if err != nil {
		return &user.LoginResponse{
			Code: user.Code_ERROR_UNKNOWN,
		}, err
	}

	return &user.LoginResponse{
		Code: user.Code_OK,
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
