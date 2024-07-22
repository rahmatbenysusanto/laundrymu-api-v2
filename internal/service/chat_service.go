package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"laundrymu-api/internal/repository"
)

type ChatService interface {
	FindByTokoId(ctx *fiber.Ctx, tokoId int) ([]*entity.Chat, error)
	Create(ctx *fiber.Ctx, chat *entity.ChatCreate) error
}

type chatService struct {
	DB             *sql.DB
	ChatRepository repository.ChatRepository
}

func NewChatService(DB *sql.DB) ChatService {
	return &chatService{
		DB:             DB,
		ChatRepository: repository.NewChatRepository(),
	}
}

func (c *chatService) FindByTokoId(ctx *fiber.Ctx, tokoId int) ([]*entity.Chat, error) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}

	chat, err := c.ChatRepository.FindByTokoId(ctx, tx, tokoId)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	} else {
		_ = tx.Commit()
		return chat, nil
	}
}

func (c *chatService) Create(ctx *fiber.Ctx, chat *entity.ChatCreate) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}

	err = c.ChatRepository.Create(ctx, tx, chat)

	if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		_ = tx.Commit()
		return nil
	}
}
