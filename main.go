package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/davidlozano107/twitter-golang/awsgo"
	"github.com/davidlozano107/twitter-golang/bd"
	"github.com/davidlozano107/twitter-golang/handlers"
	"github.com/davidlozano107/twitter-golang/models"
	"github.com/davidlozano107/twitter-golang/secretmanager"
	"os"
	"strings"
)

func main() {
	lambda.Start(EjectorLambda)
}

func EjectorLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()

	if !ParameterValidation() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Debe incluir 'SecretName','BucketName','UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitter-go"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), os.Getenv(path))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)

	//Conexión BD

	if err = bd.ConectarDB(awsgo.Ctx); err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectado la base de datos" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	resAPI := handlers.Manejadores(awsgo.Ctx, request)
	if resAPI.CustomResp != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resAPI.Status,
			Body:       resAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return resAPI.CustomResp, nil
	}
}

func ParameterValidation() bool {
	secretsName := []string{"SecretName", "BucketName", "UrlPrefix"}
	for _, secretName := range secretsName {
		if !validateIfExist(secretName) {
			return false
		}
	}
	return true
}

func validateIfExist(key string) bool {
	_, traeParametro := os.LookupEnv(key)
	if !traeParametro {
		return false
	}
	return true
}
