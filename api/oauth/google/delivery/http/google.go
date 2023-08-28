package http

import (
	"net/http"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	"github.com/gorilla/mux"
)

type googleHandler struct {
	googleUsecase usecase.GoogleUsecase
}

func NewGoogleHandler(r *mux.Router, g usecase.GoogleUsecase) {
	handler := &googleHandler{g}

	oauth := r.PathPrefix("/oauth/google").Subrouter()
	oauth.HandleFunc("/redirect", handler.redirectToGoogle)
}

func (g *googleHandler) redirectToGoogle(w http.ResponseWriter, r *http.Request) {

}
