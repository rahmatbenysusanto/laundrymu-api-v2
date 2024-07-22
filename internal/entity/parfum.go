package entity

import "time"

type Parfum struct {
	Id        int       `json:"id"`
	TokoId    int       `json:"toko_id"`
	Nama      string    `json:"nama"`
	Harga     int       `json:"harga"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ParfumToko struct {
	TokoId int `json:"toko_id" validate:"required"`
}

type ParfumId struct {
	Id int `json:"id" validate:"required"`
}

type ParfumCreate struct {
	TokoId int    `json:"toko_id" validate:"required"`
	Nama   string `json:"nama" validate:"required, min=1,max=255"`
	Harga  int    `json:"harga" validate:"required"`
}

type ParfumUpdate struct {
	Id    int    `json:"id" validate:"required"`
	Nama  string `json:"nama" validate:"min=1,max=255"`
	Harga int    `json:"harga" validate:"min=1"`
}
