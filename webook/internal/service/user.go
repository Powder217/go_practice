package service

import (
	"context"
	"errors"

	"github.com/Powder217/go_practice/webook/internal/domain"
	"github.com/Powder217/go_practice/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var(
	ErrUserDuplicateEmail=repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码错误")
)

type UserService struct{
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository)*UserService{
	return 	&UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context,u domain.User) error{
	hash,err:=bcrypt.GenerateFromPassword([]byte(u.PassWord),bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	u.PassWord=string(hash)

	return svc.repo.Create(ctx,u)
}

func (svc *UserService) Login(ctx context.Context,email,password string) error{
	u,err:=svc.repo.FindByEmail(ctx,email)

	if err==gorm.ErrRecordNotFound{
		return ErrInvalidUserOrPassword
	}
	if err!=nil{
		return err
	}

	err=bcrypt.CompareHashAndPassword([]byte(u.PassWord),[]byte(password))
	if err!=nil{
		return ErrInvalidUserOrPassword
	}
	return nil
}