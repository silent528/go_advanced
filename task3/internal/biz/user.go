package biz

import "context"

type User struct {
	Uid      int64
	UserName string
}

type UserRepo interface {
	GetUserInfo(ctx context.Context, uid int64) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc UserUsecase) GetUserName(ctx context.Context, uid int64) (string, error) {
	userInfo, err := uc.repo.GetUserInfo(ctx, uid)
	if err != nil {
		return "", err
	}
	return userInfo.UserName, nil
}
