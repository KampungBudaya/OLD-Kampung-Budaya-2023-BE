package repository

import (
	"context"
	"fmt"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/jmoiron/sqlx"
)

type FordaRepositoryImpl interface {
	Create(req domain.FordaRegister, ctx context.Context) (int64, error)
	FindByEmail(email string, ctx context.Context) (*domain.Forda, error)
	CreatePhotoPayment(link string, id int, ctx context.Context) (string, error)
}

type FordaRepository struct {
	mysql *sqlx.DB
}

func NewFordaRepository(mysql *sqlx.DB) FordaRepositoryImpl {
	return &FordaRepository{
		mysql: mysql,
	}
}

func (r *FordaRepository) Create(req domain.FordaRegister, ctx context.Context) (int64, error) {
	res, err := r.mysql.Exec(queryCreateForda, req.Name, req.Email, req.Phone)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *FordaRepository) FindByEmail(email string, ctx context.Context) (*domain.Forda, error) {
	var forda domain.FordaDB
	queryFindByEmail := fmt.Sprintf(queryFindForda, "WHERE email = ?")
	if err := r.mysql.QueryRowx(queryFindByEmail, email).StructScan(&forda); err != nil {
		return nil, err
	}
	return forda.Formatting(), nil
}

func (r *FordaRepository) CreatePhotoPayment(link string, id int, ctx context.Context) (string, error) {
	_, err := r.mysql.Exec(queryCreatePhoto, link, id)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (r *FordaRepository) FindByID(id int, ctx context.Context) (*domain.Forda, error) {
	var forda domain.FordaDB
	queryFindByID := fmt.Sprintf(queryFindForda, "WHERE id = ?")
	if err := r.mysql.QueryRowx(queryFindByID, id).StructScan(&forda); err != nil {
		return nil, err
	}
	return forda.Formatting(), nil
}
