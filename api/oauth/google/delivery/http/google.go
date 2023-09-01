package http

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type googleHandler struct {
	googleUsecase usecase.GoogleUsecase
	googleOAuth   *oauth2.Config
}

func NewGoogleHandler(r *mux.Router, g usecase.GoogleUsecase, o *oauth2.Config) {
	handler := &googleHandler{
		googleUsecase: g,
		googleOAuth:   o,
	}

	oauth := r.PathPrefix("/oauth/google").Subrouter()
	oauth.HandleFunc("/redirect", handler.redirect).Methods(http.MethodGet)
	oauth.HandleFunc("/callback", handler.handleCallback).Methods(http.MethodGet)
}

func (g *googleHandler) redirect(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	redirectUrl := g.googleOAuth.AuthCodeURL(state)

	http.SetCookie(w, &http.Cookie{Name: "state", Value: state})
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func (g *googleHandler) handleCallback(w http.ResponseWriter, r *http.Request) {
	stateCookie, err := r.Cookie("state")
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "State cookie does not exist!")
		return
	}

	stateForm := r.FormValue("state")

	if stateCookie.Value != stateForm {
		response.Fail(w, http.StatusBadRequest, "State cookie has been tampered!")
		return
	}

	code := r.FormValue("code")
	token, err := g.googleOAuth.Exchange(context.Background(), code)
	if err != nil {
		response.Fail(w, http.StatusUnauthorized, "Failed to exchange authorization token")
	}

	oauthGoogleUrlAPI := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	client := g.googleOAuth.Client(context.Background(), token)
	resp, err := client.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		response.Fail(w, http.StatusForbidden, "Code-Token Resource Exchange Failed")
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, "Fail to read resource body")
		return
	}

	var user domain.GoogleUser

	if err := json.Unmarshal(data, &user); err != nil {
		response.Fail(w, http.StatusInternalServerError, "Fail to Unmarshal User Data")
		return
	}

	if err := g.googleUsecase.Find(user.Id); err != nil {
		response.Fail(w, http.StatusUnauthorized, "User does NOT exists in the system")
		return
	}
}
