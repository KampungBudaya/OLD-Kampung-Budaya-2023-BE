package model

import (
	"errors"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/validator"
)

type User struct {
	ID           int       `json:"id"`
	Provider     string    `json:"provider"`
	ProviderID   string    `json:"provider_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Phone        string    `json:"phone"`
	PaymentPhoto string    `json:"link_photo"`
	CreateAt     time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"updated_at"`
}

type UserRegister struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func (f *UserRegister) Validate() error {
	if f.Name == nil || f.Email == nil || f.Phone == nil {
		return errors.New("FIELD CAN'T BE EMPTY")
	}
	if err := validator.ValidateEmail(*f.Email); err != nil {
		return err
	}
	if err := validator.ValidatePhone(*f.Phone); err != nil {
		return err
	}
	return nil
}

type UserLogin struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (f *UserLogin) Validate() error {
	if f.Email == nil || f.Password == nil {
		return errors.New("FIELD CAN'T BE EMPTY")
	}
	if err := validator.ValidateEmail(*f.Email); err != nil {
		return err
	}
	return nil
}

type UserDB struct {
	ID           int       `db:"id"`
	Provider     *string   `db:"provider"`
	ProviderID   *string   `db:"provider_id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Password     *string   `db:"password"`
	Phone        string    `db:"phone"`
	PaymentPhoto string    `db:"link_photo"`
	CreateAt     time.Time `db:"created_at"`
	UpdateAt     time.Time `db:"updated_at"`
}

func (fdb *UserDB) Formatting() *User {
	user := User{
		ID:           fdb.ID,
		Name:         fdb.Name,
		Email:        fdb.Email,
		Phone:        fdb.Phone,
		PaymentPhoto: fdb.PaymentPhoto,
		CreateAt:     fdb.CreateAt,
		UpdateAt:     fdb.UpdateAt,
	}

	if fdb.Provider != nil {
		user.Provider = *fdb.Provider
	}

	if fdb.ProviderID != nil {
		user.ProviderID = *fdb.ProviderID
	}

	if fdb.Password != nil {
		user.Password = *fdb.Password
	}
	return &user
}
