package bd

import (
	"context"
	"github.com/davidlozano107/twitter-golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	db := MongoCN.Database(DatabaseName)
	collection := db.Collection("usuarios")

	condition := bson.M{"email": email}
	var result models.Usuario

	id := result.Id.Hex()

	if err := collection.FindOne(context.TODO(), condition).Decode(&result); err != nil {
		return result, false, id
	}

	return result, true, id
}
