package model

import "time"

type User struct {
	Id         int
	Provider   string
	ProviderId string
	Username   string
	Email      string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
