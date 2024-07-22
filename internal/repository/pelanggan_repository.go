package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"time"
)

type PelangganRepository interface {
	FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Pelanggan, error)
	Create(ctx *fiber.Ctx, tx *sql.Tx, pelanggan *entity.PelangganRequest) error
}

type pelangganRepository struct{}

func NewPelangganRepository() PelangganRepository {
	return &pelangganRepository{}
}

func (p *pelangganRepository) FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, tokoId int) ([]*entity.Pelanggan, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM pelanggan WHERE toko_id = ?", tokoId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var results []*entity.Pelanggan
	for rows.Next() {
		var pelanggan entity.Pelanggan
		err := rows.Scan(
			&pelanggan.Id,
			&pelanggan.TokoId,
			&pelanggan.Nama,
			&pelanggan.NoHp,
			&pelanggan.CreatedAt,
			&pelanggan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &pelanggan)
	}
	return results, nil
}

func (p *pelangganRepository) Create(ctx *fiber.Ctx, tx *sql.Tx, pelanggan *entity.PelangganRequest) error {
	SQL := "INSERT INTO pelanggan (toko_id, nama, no_hp, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx.Context(), SQL, pelanggan.TokoId, pelanggan.Nama, pelanggan.NoHp, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
