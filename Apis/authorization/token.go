package authorization

import (
	"Api-Go/pkg/model"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "NCG",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("Token no valido")
	}
	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("No se pudo obtener los claim")
	}
	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
