package entity

import "time"

type Diskon struct {
	Id        int       `json:"id"`
	TokoId    int       `json:"toko_id"`
	Nama      string    `json:"nama"`
	Type      string    `json:"type"`
	Nominal   int       `json:"nominal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DiskonToko struct {
	TokoId int `json:"toko_id" validate:"required"`
}

type DiskonId struct {
	Id int `json:"id" validate:"required"`
}

type DiskonCreate struct {
	TokoId  int    `json:"toko_id" validate:"required"`
	Nama    string `json:"nama" validate:"required,min=1,max=255"`
	Type    string `json:"type" validate:"required"`
	Nominal int    `json:"nominal" validate:"required"`
}
