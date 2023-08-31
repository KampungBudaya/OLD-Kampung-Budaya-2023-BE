package middleware

import (
	"net/http"
	"strings"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/jwt"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
)

func ValidateJWT() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			response.Fail(&w, http.StatusUnauthorized, "TOKEN NOT FOUND")
			return
		}

		tokens := strings.SplitN(token, " ", 2)
		if len(tokens) != 2 {
			response.Fail(&w, http.StatusUnauthorized, "INVALID TOKEN")
			return
		}

		token = tokens[1]
		claims, err := jwt.ParseJWT(token)
		if err != nil {
			response.Fail(&w, http.StatusUnauthorized, err.Error())
			return
		}

		id := claims["id"].(string)
		roles := claims["roles"]
		rolesConcat := strings.Join(roles.([]string), " ")

		r.Header.Set("id", id)
		r.Header.Set("role", rolesConcat)
	}
}
