package encryption

import (
	"github.com/MauricioGZ/Cubes/internal/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})
	return token.SignedString([]byte(key))
}

func ParseLoginJWT(_token string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(_token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
