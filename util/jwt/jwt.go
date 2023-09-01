package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/domain"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user *domain.Forda, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"roles": []string{role},
	})

	signedToken, err := token.SignedString([]byte("SECRET KEY"))
	if err != nil {
		return "", errors.New("FAILED TO GENERATE JWT")
	}

	return signedToken, nil
}

func ParseJWT(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET KEY")), nil
	})

	if err != nil {
		return nil, errors.New("FAILED TO PARSE JWT")
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !decoded.Valid || !ok {
		return nil, errors.New("INVALID JWT")
	}
	return claims, nil
}
