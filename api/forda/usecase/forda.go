package usecase

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/database"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/model"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/jwt"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/password"
)

type FordaUsecaseImpl interface {
	Register(req model.FordaRegister, ctx context.Context) (int64, error)
	Login(req model.FordaLogin, ctx context.Context) (string, error)
	UploadPhotoPayment(file *multipart.FileHeader, id int, ctx context.Context) (string, error)
}

type FordaUsecase struct {
	FordaRepo repository.FordaRepositoryImpl
}

func NewFordaUsecase(fordaRepo repository.FordaRepositoryImpl) FordaUsecaseImpl {
	return &FordaUsecase{
		FordaRepo: fordaRepo,
	}
}

func (u *FordaUsecase) Register(req model.FordaRegister, ctx context.Context) (int64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	id, err := u.FordaRepo.Create(req, ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *FordaUsecase) Login(req model.FordaLogin, ctx context.Context) (string, error) {
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

func (u *FordaUsecase) UploadPhotoPayment(file *multipart.FileHeader, id int, ctx context.Context) (string, error) {
	var (
		link string
		err  error
	)

	if file == nil {
		link = ""
	} else {
		supClient := database.NewSupabaseClient()
		link, err = supClient.UploadFile(file)
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
