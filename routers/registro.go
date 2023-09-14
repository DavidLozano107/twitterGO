package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davidlozano107/twitter-golang/bd"
	"github.com/davidlozano107/twitter-golang/models"
)

func Registro(ctx context.Context) models.ResApi {
	var t models.Usuario
	var r models.ResApi
	r.Status = 400
	fmt.Println("Entre a registro")

	body := ctx.Value(models.Key("body")).(string)

	if err := json.Unmarshal([]byte(body), &t); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 carecteres."
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese email."
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario. " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println(r.Message)
	return r
}
