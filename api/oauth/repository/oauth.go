package repository

import (
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/model"
	"github.com/jmoiron/sqlx"
)

type OAuthRepository interface {
	Store(user *model.User) error
	GetAll(pageNum int) ([]model.User, model.PaginateMeta, error)
	GetByEmail(email string) (model.User, error)
	GetByProviderId(id string) (model.User, error)
	Update(user *model.User) error
}

type oauthRepository struct {
	conn *sqlx.DB
}

func NewOAuthRepository(conn *sqlx.DB) OAuthRepository {
	return &oauthRepository{conn}
}

func (o *oauthRepository) Store(user *model.User) error {
	stmt := "INSERT INTO users (provider, provider_id, username, email) VALUES (?, ?, ?, ?)"

	if _, err := o.conn.Exec(stmt, user.Provider, user.ProviderId, user.Username, user.Email); err != nil {
		return err
	}

	return nil
}

func (o *oauthRepository) GetAll(pageNum int) ([]model.User, model.PaginateMeta, error) {
	var users []model.User

	limit := 20
	cursor := pageNum * limit

	if err := o.conn.Select(&users, "SELECT * FROM users WHERE id <= ? ORDER BY id ASC LIMIT ?", cursor, limit); err != nil {
		return nil, model.PaginateMeta{}, err
	}

	var meta model.PaginateMeta

	if err := o.conn.Select(&meta.Count, "SELECT COUNT(*) FROM users"); err != nil {
		return nil, model.PaginateMeta{}, err
	}

	meta.Starts = cursor - 19
	meta.Ends = cursor

	return users, meta, nil
}

func (o *oauthRepository) GetByEmail(email string) (model.User, error) {
	var user model.User

	if err := o.conn.Get(&user, "SELECT * FROM users WHERE email = ?", email); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (o *oauthRepository) GetByProviderId(id string) (model.User, error) {
	var user model.User

	if err := o.conn.Get(&user, "SELECT * FROM users WHERE provider_id = ?", id); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (o *oauthRepository) Update(user *model.User) error {
	if _, err := o.conn.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Id); err != nil {
		return err
	}

	return nil
}
