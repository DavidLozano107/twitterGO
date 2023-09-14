package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	Id              primitive.ObjectID `bson:"_id", json:"id",omitempty`
	Nombre          string             `bson:"nombre", json:"nombre", omitempty`
	Apellido        string             `bson:"apellido", json:apellido, omitempty`
	FechaNacimiento string             `bson:"fecha_nacimiento", json:"fecha_nacimiento", omitempty`
	Email           string             `bson:"email", json:"email"`
	Password        string             `bson:"password", json:"password", omitempty`
	Avtar           string             `bson:"avtar", json:"avtar", omitempty`
	Banner          string             `bson:"banner", json:"banner", omitempty`
	Biografia       string             `bson:"biografia", json:"biografia", omitempty`
	Ubicacion       string             `bson:"ubicacion", json:"ubicacion", ubicacion`
	SitioWeb        string             `bson:"sitio_web", json:"sitio_web", omitempty`
}
