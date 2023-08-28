package repository

import (
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/jmoiron/sqlx"
)

type OAuthRepository interface {
	Store(user *domain.ProviderUserRegister) error
	GetByEmail(email string) (domain.User, error)
	GetByProviderId(id string) (domain.User, error)
}

type oauthRepository struct {
	conn *sqlx.DB
}

func NewOAuthRepository(conn *sqlx.DB) OAuthRepository {
	return &oauthRepository{conn}
}

func (o *oauthRepository) Store(user *domain.ProviderUserRegister) error {
	stmt := "INSERT INTO users (provider, provider_id, username, email) VALUES (?, ?, ?, ?)"

	if _, err := o.conn.Exec(stmt, user.Provider, user.ProviderId, user.Username, user.Email); err != nil {
		return err
	}

	return nil
}

func (o *oauthRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := o.conn.Get(&user, "SELECT * FROM users WHERE email = ?", email); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (o *oauthRepository) GetByProviderId(id string) (domain.User, error) {
	var user domain.User

	if err := o.conn.Get(&user, "SELECT * FROM users WHERE provider_id = ?", id); err != nil {
		return domain.User{}, err
	}

	return user, nil
}
