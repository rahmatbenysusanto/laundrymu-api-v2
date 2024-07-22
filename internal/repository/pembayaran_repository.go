package repository

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"time"
)

type PembayaranRepository interface {
	FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Pembayaran, error)
	FindById(ctx *fiber.Ctx, tx *sql.Tx, id int) (*entity.Pembayaran, error)
	Create(ctx *fiber.Ctx, tx *sql.Tx, pembayaran *entity.PembayaranCreate) error
}

type pembayaranRepository struct{}

func NewPembayaranRepository() PembayaranRepository {
	return &pembayaranRepository{}
}

func (p *pembayaranRepository) FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Pembayaran, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM pembayaran_toko WHERE toko_id = ?", tokoId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var results []*entity.Pembayaran
	for rows.Next() {
		var result entity.Pembayaran
		err := rows.Scan(
			&result.Id,
			&result.TokoId,
			&result.Nama,
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

func (p *pembayaranRepository) FindById(ctx *fiber.Ctx, tx *sql.Tx, id int) (*entity.Pembayaran, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM pembayaran_toko WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var result entity.Pembayaran
	if rows.Next() {
		err := rows.Scan(
			&result.Id,
			&result.TokoId,
			&result.Nama,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (p *pembayaranRepository) Create(ctx *fiber.Ctx, tx *sql.Tx, pembayaran *entity.PembayaranCreate) error {
	SQL := "INSERT INTO pembayaran (toko_id, nama, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx.Context(), SQL, pembayaran.TokoId, pembayaran.Nama, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
