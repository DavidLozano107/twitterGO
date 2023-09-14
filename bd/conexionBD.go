package bd

import (
	"context"
	"fmt"
	"github.com/davidlozano107/twitter-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectarDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majoriry", user, passwd, host)

	var clientsOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientsOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexión Exitosa con la DB")
	MongoCN = client

	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
