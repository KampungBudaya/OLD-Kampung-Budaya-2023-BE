package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/model"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
)

const reqTO = "REQUEST TIMEOUT"

type FordaHandler struct {
	Router       *mux.Router
	FordaUsecase usecase.FordaUsecaseImpl
}

func NewFordaHandler(router *mux.Router, fordaUsecase usecase.FordaUsecaseImpl) {
	handler := &FordaHandler{
		Router:       router,
		FordaUsecase: fordaUsecase,
	}
	handler.Router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	handler.Router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	handler.Router.HandleFunc("/upload-photo-payment/{id}", handler.UploadPhotoPayment).Methods(http.MethodPost)
}

func (h *FordaHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var (
		err  error
		code = http.StatusOK
		data interface{}
	)

	defer func() {
		if err != nil {
			response.Fail(&w, code, err.Error())
			return
		}
		response.Success(&w, code, data)
	}()

	var request model.FordaRegister

	body, err := io.ReadAll(r.Body)
	if err != nil {
		code = http.StatusBadRequest
		err = errors.New("FAILED TO READ REQUEST BODY")
		return
	}

	if err = json.Unmarshal(body, &request); err != nil {
		err = errors.New("FAILED TO UNMARSHAL REQUEST BODY")
		code = http.StatusBadRequest
		return
	}

	forda, err := h.FordaUsecase.Register(request, ctx)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-ctx.Done():
		code = http.StatusRequestTimeout
		err = errors.New(reqTO)
		return
	default:
		data = forda
	}
}

func (h *FordaHandler) UploadPhotoPayment(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var (
		err  error
		code = http.StatusOK
		data interface{}
	)

	defer func() {
		if err != nil {
			response.Fail(&w, code, err.Error())
			return
		}
		response.Success(&w, code, data)
	}()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = errors.New("FAILED TO CONVERT STRING TO INTEGER")
		code = http.StatusBadRequest
		return
	}

	var file *multipart.FileHeader
	_, file, err = r.FormFile("photo-payment")
	if err != nil {
		file = nil
	}

	link, err := h.FordaUsecase.UploadPhotoPayment(file, id, ctx)
	if err != nil {
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-ctx.Done():
		code = http.StatusRequestTimeout
		err = errors.New(reqTO)
		return
	default:
		data = link
	}
}

func (h *FordaHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	var (
		err  error
		code = http.StatusOK
		data interface{}
	)

	defer func() {
		if err != nil {
			response.Fail(&w, code, err.Error())
			return
		}
		response.Success(&w, code, data)
	}()

	var request model.FordaLogin
	body, err := io.ReadAll(r.Body)
	if err != nil {
		code = http.StatusBadRequest
		err = errors.New("FAILED TO READ REQUEST BODY")
		return
	}

	if err = json.Unmarshal(body, &request); err != nil {
		err = errors.New("FAILED TO UNMARSHAL REQUEST BODY")
		code = http.StatusBadRequest
		return
	}

	forda, err := h.FordaUsecase.Login(request, ctx)

	select {
	case <-ctx.Done():
		code = http.StatusRequestTimeout
		err = errors.New(reqTO)
		return
	default:
		data = forda
	}
}
