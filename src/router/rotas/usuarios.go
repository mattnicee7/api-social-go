package rotas

import (
	"api-social-go/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:              "/usuarios",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuarios,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletarUsuario,
		RequerAutenticao: true,
	},
}
