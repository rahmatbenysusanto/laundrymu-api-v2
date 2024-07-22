package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
)

type UserRepository interface {
	Login(ctx *fiber.Ctx, tx *sql.Tx, username string, password string) (*entity.User, error)
	Create(ctx *fiber.Ctx, tx *sql.Tx, user *entity.UserRequest) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) Login(ctx *fiber.Ctx, tx *sql.Tx, username string, password string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Create(ctx *fiber.Ctx, tx *sql.Tx, user *entity.UserRequest) error {
	//TODO implement me
	panic("implement me")
}
