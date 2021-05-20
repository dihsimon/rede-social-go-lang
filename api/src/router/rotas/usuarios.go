package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:               "/usuarios",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarUsuario,
		RequerAutenticaco: false,
	},
	{
		URI:               "/usuarios",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuarios,
		RequerAutenticaco: false,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuario,
		RequerAutenticaco: false,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodPut,
		Funcao:            controllers.AtualizarUsuario,
		RequerAutenticaco: false,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarUsuario,
		RequerAutenticaco: false,
	},
}
