package entity

import "time"

type Pembayaran struct {
	Id        int       `json:"id"`
	TokoId    int       `json:"toko_id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PembayaranToko struct {
	TokoId int `json:"toko_id" validate:"required"`
}

type PembayaranId struct {
	Id int `json:"id" validate:"required"`
}

type PembayaranCreate struct {
	TokoId int    `json:"toko_id" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
}
