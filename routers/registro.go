package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jhonatan1085/redsocial/bd"
	"github.com/jhonatan1085/redsocial/models"
)

/*Registro es la funcion ppara crear en la bd el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Email ya existe", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un erroral intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
