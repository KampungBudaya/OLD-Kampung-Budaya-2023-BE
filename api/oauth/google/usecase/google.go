package usecase

import (
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
)

type GoogleUsecase interface {
	Find(id string) error
	Register(user domain.ProviderUserRegister) error
}

type googleUsecase struct {
	oauthRepository repository.OAuthRepository
}

func NewGoogleUsecase(o repository.OAuthRepository) GoogleUsecase {
	return &googleUsecase{o}
}

func (g *googleUsecase) Find(id string) error {

	if _, err := g.oauthRepository.GetByProviderId(id); err != nil {
		return err
	}

	return nil
}

func (g *googleUsecase) Register(user domain.ProviderUserRegister) error {
	if err := g.oauthRepository.Store(&user); err != nil {
		return err
	}

	return nil
}
