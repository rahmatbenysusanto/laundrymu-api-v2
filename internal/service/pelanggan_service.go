package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"laundrymu-api/internal/entity"
	"laundrymu-api/internal/repository"
)

type PelangganService interface {
	FindByTokoId(ctx *fiber.Ctx, tokoId int) ([]*entity.Pelanggan, error)
	Create(ctx *fiber.Ctx, pelanggan *entity.PelangganRequest) error
}

type pelangganService struct {
	DB                  *sql.DB
	PelangganRepository repository.PelangganRepository
}

func NewPelangganService(DB *sql.DB) PelangganService {
	return &pelangganService{
		DB:                  DB,
		PelangganRepository: repository.NewPelangganRepository(),
	}
}

func (s *pelangganService) FindByTokoId(ctx *fiber.Ctx, tokoId int) ([]*entity.Pelanggan, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	result, err := s.PelangganRepository.FindByTokoId(ctx, tx, tokoId)
	if err != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}

	return result, nil
}

func (s *pelangganService) Create(ctx *fiber.Ctx, pelanggan *entity.PelangganRequest) error {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	err = s.PelangganRepository.Create(ctx, tx, pelanggan)
	if err != nil {
		_ = tx.Rollback()
		return nil
	} else {
		_ = tx.Commit()
		return err
	}
}
