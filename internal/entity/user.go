package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Nama      string    `json:"nama"`
	NoHp      string    `json:"no_hp"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Otp       int       `json:"otp"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Nama     string `json:"nama" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	Token        string      `json:"token"`
	RefreshToken string      `json:"refreshToken"`
	User         interface{} `json:"user"`
}
