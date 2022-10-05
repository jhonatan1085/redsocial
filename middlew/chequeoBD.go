package middlew

import (
	"net/http"

	"github.com/jhonatan1085/redsocial/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
