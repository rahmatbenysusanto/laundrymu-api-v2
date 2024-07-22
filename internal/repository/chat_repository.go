package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"time"
)

type ChatRepository interface {
	FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Chat, error)
	Create(ctx *fiber.Ctx, tx *sql.Tx, chat *entity.ChatCreate) error
}

type chatRepository struct{}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

func (c *chatRepository) FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Chat, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM chat WHERE toko_id=?", tokoId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var chats []*entity.Chat
	for rows.Next() {
		var chat = entity.Chat{}
		err := rows.Scan(
			&chat.Id,
			&chat.TokoId,
			&chat.Role,
			&chat.Chat,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}
	return chats, nil
}

func (c *chatRepository) Create(ctx *fiber.Ctx, tx *sql.Tx, chat *entity.ChatCreate) error {
	SQL := "INSERT INTO chat (toko_id, role, chat, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx.Context(), SQL, chat.TokoId, chat.Role, chat.Chat, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
