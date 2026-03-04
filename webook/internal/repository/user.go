package repository

import (
	"context"

	"github.com/Powder217/go_practice/webook/internal/domain"
	"github.com/Powder217/go_practice/webook/internal/repository/dao"
)

var(
	ErrUserDuplicateEmail=dao.ErrUserDuplicateEmail
)

type UserRepository struct{
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO)UserRepository{
	return UserRepository{
		dao: dao,
	}
}
func (r *UserRepository) Create (ctx context.Context,u domain.User)error{
	return r.dao.Insert(ctx,dao.User{
		Email: u.Email,
		PassWord: u.Password,
	})
}

