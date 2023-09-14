package bd

import (
	"context"
	"github.com/davidlozano107/twitter-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(user models.Usuario) (string, bool, error) {
	db := MongoCN.Database(DatabaseName)
	collection := db.Collection("usuarios")

	user.Password, _ = EncriptarPassword(user.Password)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
