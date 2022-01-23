package data

import (
	"context"
	"go_advanced/graduation/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

type userRepo struct {
	data *Data
}

func (repo *userRepo) GetUserInfo(ctx context.Context, uid int64) (u *biz.User, err error) {
	userRow, err := repo.data.db.User.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Uid:      userRow.UID,
		UserName: userRow.Username,
	}, nil
}
