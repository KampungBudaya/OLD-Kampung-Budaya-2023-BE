package domain

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

type ProviderUserRegister struct {
	Provider   string
	ProviderId string
	Username   string
	Email      string
}

type GoogleUser struct {
	Id            string
	Email         string
	VerifiedEmail bool
	Name          string
	GivenName     string
	Picture       string
	Locale        string
}
