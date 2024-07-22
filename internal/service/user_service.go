package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
)

type UserService interface {
	Login(ctx *fiber.Ctx, user *entity.UserLogin) (*entity.UserLoginResponse, error)
	Register(ctx *fiber.Ctx, user *entity.UserRequest) error
}

type userService struct {
	DB *sql.DB
}

func NewUserService(DB *sql.DB) UserService {
	return &userService{
		DB: DB,
	}
}

func (u *userService) Login(ctx *fiber.Ctx, user *entity.UserLogin) (*entity.UserLoginResponse, error) {
	token, _ := CreateToken("tes", "tes", "tes", "tes")
	result := entity.UserLoginResponse{
		Token:        token,
		RefreshToken: "12345",
		User:         "hahaha",
	}
	return &result, nil
}

func (u *userService) Register(ctx *fiber.Ctx, user *entity.UserRequest) error {
	//TODO implement me
	panic("implement me")
}
