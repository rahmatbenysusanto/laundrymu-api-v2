package entity

import "time"

type Pelanggan struct {
	Id        int       `json:"id"`
	TokoId    int       `json:"toko_id"`
	Nama      string    `json:"nama"`
	NoHp      string    `json:"no_hp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PelangganRequest struct {
	TokoId int    `json:"toko_id" validate:"required"`
	Nama   string `json:"nama" validate:"required,min=1,max=100"`
	NoHp   string `json:"no_hp" validate:"required,min=1,max=15"`
}

type PelangganByTokoId struct {
	TokoId int `json:"toko_id" validate:"required"`
}
