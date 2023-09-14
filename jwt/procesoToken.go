package jwt

import (
	"errors"
	"strings"

	"github.com/davidlozano107/twitter-golang/models"
	"github.com/golang-jwt/jwt/v5"
)

var Email string
var IdUsuario string

func ProcesoToken(tk string, JWTSing string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSing)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")

 	if len(splitToken) != 2 {
		 return &claims, false, "", errors.New("Formato de token inválido.")
	}
	
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	
	if err == nil {
		//rutina que chequea contra la BD
	}
	
	if !tkn.Valid {
		return &claims, false, "", errors.New("Token Inválido")
	}

	return &claims, false, "", nil

}
