package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jhonatan1085/redsocial/bd"
	"github.com/jhonatan1085/redsocial/jwt"
	"github.com/jhonatan1085/redsocial/models"
)

/*Login realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}

	jwtkey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al generar el tokken correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated) //devuelve estado 200
	json.NewEncoder(w).Encode(resp)

	//grabar una kokie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
