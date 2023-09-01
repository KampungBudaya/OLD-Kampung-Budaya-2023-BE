package usecase

import (
	"context"
	"fmt"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/config"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/jwt"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/password"
)

type FordaUsecaseImpl interface {
	Register(req domain.FordaRegister, ctx context.Context) (int64, error)
	Login(req domain.FordaLogin, ctx context.Context) (string, error)
	UploadPhotoPayment(fileBytes []byte, id int, ctx context.Context) (string, error)
}

type FordaUsecase struct {
	FordaRepo repository.FordaRepositoryImpl
}

func NewFordaUsecase(fordaRepo repository.FordaRepositoryImpl) FordaUsecaseImpl {
	return &FordaUsecase{
		FordaRepo: fordaRepo,
	}
}

func (u *FordaUsecase) Register(req domain.FordaRegister, ctx context.Context) (int64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	id, err := u.FordaRepo.Create(req, ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *FordaUsecase) Login(req domain.FordaLogin, ctx context.Context) (string, error) {
	forda, err := u.FordaRepo.FindByEmail(*req.Email, ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = password.ComparePassword(*req.Password, forda.Password)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateJWT(forda, "forda")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *FordaUsecase) UploadPhotoPayment(fileBytes []byte, id int, ctx context.Context) (string, error) {
	var (
		link string
		err  error
	)

	if fileBytes == nil {
		link = ""
	} else {
		fb := config.InitFirebase()
		link, err = fb.UploadFile(ctx, fileBytes, fmt.Sprintf("photo-payment-%d", id))
		if err != nil {
			return "", err
		}
	}

	link, err = u.FordaRepo.CreatePhotoPayment(link, id, ctx)
	if err != nil {
		return "", err
	}
	return link, nil
}
