package main

import (
	"log"

	"github.com/jhonatan1085/redsocial/bd"
	"github.com/jhonatan1085/redsocial/handlers"
)

func main() {
	if bd.ChequeConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
