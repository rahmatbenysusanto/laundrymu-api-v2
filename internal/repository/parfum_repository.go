package repository

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"time"
)

type ParfumRepository interface {
	FindById(ctx *fiber.Ctx, tx *sql.Tx, Id int) (*entity.Parfum, error)
	FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, TokoId int) ([]*entity.Parfum, error)
	Create(ctx *fiber.Ctx, tx *sql.Tx, Parfum *entity.ParfumCreate) error
	Update(ctx *fiber.Ctx, tx *sql.Tx, Parfum *entity.ParfumUpdate) error
	Delete(ctx *fiber.Ctx, tx *sql.Tx, Id int) error
}

type parfumRepository struct{}

func NewParfumRepository() ParfumRepository {
	return &parfumRepository{}
}

func (p *parfumRepository) FindById(ctx *fiber.Ctx, tx *sql.Tx, Id int) (*entity.Parfum, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM parfum WHERE id = ?", Id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var parfum entity.Parfum
	if rows.Next() {
		err = rows.Scan(
			&parfum.Id,
			&parfum.TokoId,
			&parfum.Nama,
			&parfum.Harga,
			&parfum.CreatedAt,
			&parfum.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &parfum, nil
	} else {
		return nil, errors.New("parfum not found")
	}
}

func (p *parfumRepository) FindByTokoId(ctx *fiber.Ctx, tx *sql.Tx, TokoId int) ([]*entity.Parfum, error) {
	rows, err := tx.QueryContext(ctx.Context(), "SELECT * FROM parfum WHERE toko_id = ?", TokoId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var parfums []*entity.Parfum
	for rows.Next() {
		var parfum entity.Parfum
		err = rows.Scan(
			&parfum.Id,
			&parfum.TokoId,
			&parfum.Nama,
			&parfum.Harga,
			&parfum.CreatedAt,
			&parfum.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		parfums = append(parfums, &parfum)
	}
	return parfums, nil
}

func (p *parfumRepository) Create(ctx *fiber.Ctx, tx *sql.Tx, parfum *entity.ParfumCreate) error {
	SQL := "INSERT INTO parfum (toko_id, nama, harga, delete, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx.Context(), SQL, parfum.TokoId, parfum.Nama, parfum.Harga, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (p *parfumRepository) Update(ctx *fiber.Ctx, tx *sql.Tx, parfum *entity.ParfumUpdate) error {
	SQL := "UPDATE parfum SET nama=?, harga=?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx.Context(), SQL, time.Now(), parfum.Nama, parfum.Harga, time.Now(), parfum.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *parfumRepository) Delete(ctx *fiber.Ctx, tx *sql.Tx, Id int) error {
	SQL := "UPDATE parfum SET delete = 1, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx.Context(), SQL, time.Now(), Id)
	if err != nil {
		return err
	}
	return nil
}
