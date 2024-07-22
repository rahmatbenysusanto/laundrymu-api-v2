package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
)

type MetodePembayaranRepository interface {
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]*entity.MetodePembayaran, error)
}

type metodePembayaranRepository struct{}

func NewMetodePembayaranRepository() MetodePembayaranRepository {
	return &metodePembayaranRepository{}
}

func (m *metodePembayaranRepository) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]*entity.MetodePembayaran, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM metode_pembayaran")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var results []*entity.MetodePembayaran
	for rows.Next() {
		var result entity.MetodePembayaran
		err := rows.Scan(
			&result.Id,
			&result.MetodePembayaran,
			&result.Nama,
			&result.Nomor,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}
