package entity

import "time"

type Chat struct {
	Id        int       `json:"id"`
	TokoId    int       `json:"toko_id"`
	Role      string    `json:"role"`
	Chat      string    `json:"chat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatCreate struct {
	TokoId int    `json:"toko_id"`
	Role   string `json:"role"`
	Chat   string `json:"chat"`
}
