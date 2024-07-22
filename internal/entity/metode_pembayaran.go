package entity

import "time"

type MetodePembayaran struct {
	Id               int       `json:"id"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	Nama             string    `json:"nama"`
	Nomor            string    `json:"nomor"`
	Logo             string    `json:"logo"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
