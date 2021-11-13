package service

import (
	"context"
	v1 "go_advanced/task3/api/v1"
	"go_advanced/task3/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (us UserService) GetUserName(ctx context.Context, req *v1.GetUserNameRequest) (*v1.GetUserNameResponse, error) {
	username, err := us.uc.GetUserName(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserNameResponse{
		Username: username,
	}, nil
}
