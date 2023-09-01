package domain

import (
	"errors"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/validator"
)

type Forda struct {
	ID         int     `json:"id"`
	Provider   string  `json:"provider"`
	ProviderID string  `json:"provider_id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Phone      string  `json:"phone"`
	CreateAt   []uint8 `json:"created_at"`
	UpdateAt   []uint8 `json:"updated_at"`
}

type FordaRegister struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func (f *FordaRegister) Validate() error {
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

type FordaLogin struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (f *FordaLogin) Validate() error {
	if f.Email == nil || f.Password == nil {
		return errors.New("FIELD CAN'T BE EMPTY")
	}
	if err := validator.ValidateEmail(*f.Email); err != nil {
		return err
	}
	return nil
}

type FordaDB struct {
	ID         int     `db:"id"`
	Provider   *string `db:"provider"`
	ProviderID *string `db:"provider_id"`
	Name       string  `db:"name"`
	Email      string  `db:"email"`
	Password   string  `db:"password"`
	Phone      string  `db:"phone"`
	CreateAt   []uint8 `db:"created_at"`
	UpdateAt   []uint8 `db:"updated_at"`
}

func (fdb *FordaDB) Formatting() *Forda {
	forda := Forda{
		ID:       fdb.ID,
		Name:     fdb.Name,
		Email:    fdb.Email,
		Password: fdb.Password,
		Phone:    fdb.Phone,
		CreateAt: fdb.CreateAt,
		UpdateAt: fdb.UpdateAt,
	}

	if fdb.Provider != nil {
		forda.Provider = *fdb.Provider
	}

	if fdb.ProviderID != nil {
		forda.ProviderID = *fdb.ProviderID
	}
	return &forda
}
