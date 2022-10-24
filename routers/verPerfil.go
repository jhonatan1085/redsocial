package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jhonatan1085/redsocial/bd"
)

/*VerPerfil permite extraer los valores del perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated) //devuelve estado 200
	json.NewEncoder(w).Encode(perfil)
}
