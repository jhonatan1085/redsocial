package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC variable de conexion*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jhonatan1085:<Leonjhonatan100404>@twittersb.kkvzkk9.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("errorcito" + err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("errorcito1" + err.Error())
		return client
	}

	log.Println("conexion Exitosa con la BD")

	return client
}

/*ChequeConnection es el ping a la BD*/
func ChequeConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
