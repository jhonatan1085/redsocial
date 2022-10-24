package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jhonatan1085/redsocial/bd"
	"github.com/jhonatan1085/redsocial/models"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuleto del modelo, que se usara en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de Token invalido") // se crea un error
	}

	tk = strings.TrimSpace(splitToken[1]) //quita espacios

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()

		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}
	return claims, false, string(""), err

}
