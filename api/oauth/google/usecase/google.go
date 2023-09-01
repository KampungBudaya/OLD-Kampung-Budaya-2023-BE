package usecase

import (
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/jwt"
)

type GoogleUsecase interface {
	Login(id string) (string, error)
	Register(user domain.ProviderUserRegister) error
}

type googleUsecase struct {
	oauthRepository repository.OAuthRepository
}

func NewGoogleUsecase(o repository.OAuthRepository) GoogleUsecase {
	return &googleUsecase{o}
}

func (g *googleUsecase) Login(id string) (string, error) {

	user, err := g.oauthRepository.GetByProviderId(id)
	if err != nil {
		return "", err
	}

	forda := &domain.Forda{
		ID:         user.Id,
		Provider:   user.Provider,
		ProviderID: user.ProviderId,
		Name:       user.Username,
		Email:      user.Email,
		Password:   user.Password,
		CreateAt:   user.CreatedAt,
	}

	token, err := jwt.GenerateJWT(forda, "Admin")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (g *googleUsecase) Register(user domain.ProviderUserRegister) error {
	if err := g.oauthRepository.Store(&user); err != nil {
		return err
	}

	return nil
}
