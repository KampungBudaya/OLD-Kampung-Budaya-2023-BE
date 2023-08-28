package repository

import (
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/model"
	"github.com/jmoiron/sqlx"
)

type OAuthRepository interface {
	Store(user *model.ProviderUserRegister) error
	GetByEmail(email string) (model.User, error)
	GetByProviderId(id string) (model.User, error)
}

type oauthRepository struct {
	conn *sqlx.DB
}

func NewOAuthRepository(conn *sqlx.DB) OAuthRepository {
	return &oauthRepository{conn}
}

func (o *oauthRepository) Store(user *model.ProviderUserRegister) error {
	stmt := "INSERT INTO users (provider, provider_id, username, email) VALUES (?, ?, ?, ?)"

	if _, err := o.conn.Exec(stmt, user.Provider, user.ProviderId, user.Username, user.Email); err != nil {
		return err
	}

	return nil
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
