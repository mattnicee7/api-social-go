package rotas

import (
	"api-social-go/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:              "/login",
	Metodo:           http.MethodPost,
	Funcao:           controllers.Login,
	RequerAutenticao: false,
}
