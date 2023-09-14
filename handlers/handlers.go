package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/davidlozano107/twitter-golang/jwt"
	"github.com/davidlozano107/twitter-golang/models"
	"github.com/davidlozano107/twitter-golang/routers"
	"net/http"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	path := ctx.Value(models.Key("path")).(string)
	method := ctx.Value(models.Key("method")).(string)

	fmt.Println("Voy a procesar " + path + ">" + method)

	var r models.ResApi
	r.Status = 400

	isOk, statusCode, msg, _ := validoAuthorization(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch method {
	case "POST":
		switch path {
		case "registro":
			return routers.Registro(ctx)
		}
	case "GET":
		switch path {
		}
	case "PUT":
		switch path {
		}
	case "DELETE":
		switch path {
		}
	}

	r.Message = "Method invalid"
	return r
}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, http.StatusOK, "", models.Claim{}
	}

	token := request.Headers["authorization"]
	if len(token) == 0 {
		return false, http.StatusUnauthorized, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSing")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token", err.Error())
			return false, http.StatusUnauthorized, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token", msg)
			return false, http.StatusUnauthorized, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")

	return true, http.StatusOK, msg, *claim
}
