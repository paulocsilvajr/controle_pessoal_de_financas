package route

import (
	"controle_pessoal_de_financas/API/v1/config"
	"controle_pessoal_de_financas/API/v1/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

type Routes []Route

var MyRoutes = Routes{
	Route{
		"API",
		config.Rotas["API"].Tipo,
		config.Rotas["API"].Rota,
		controller.API,
		false,
	},
	Route{
		"Index",
		config.Rotas["Index"].Tipo,
		config.Rotas["Index"].Rota,
		controller.Index,
		true,
	},
	Route{
		"Login",
		config.Rotas["Login"].Tipo,
		config.Rotas["Login"].Rota,
		controller.Login,
		false,
	},
	Route{
		"TokenValido",
		config.Rotas["TokenValido"].Tipo,
		config.Rotas["TokenValido"].Rota,
		controller.TokenValido,
		true,
	},
	Route{
		"PessoaIndex",
		config.Rotas["PessoaIndex"].Tipo,
		config.Rotas["PessoaIndex"].Rota,
		controller.PessoaIndex,
		true,
	},
	Route{
		"PessoaShow",
		config.Rotas["PessoaShow"].Tipo,
		config.Rotas["PessoaShow"].Rota,
		controller.PessoaShow,
		true,
	},
	Route{
		"PessoaShowAdmin",
		config.Rotas["PessoaShowAdmin"].Tipo,
		config.Rotas["PessoaShowAdmin"].Rota,
		controller.PessoaShowAdmin,
		true,
	},
	Route{
		"PessoaCreate",
		config.Rotas["PessoaCreate"].Tipo,
		config.Rotas["PessoaCreate"].Rota,
		controller.PessoaCreate,
		true,
	},
	Route{
		"PessoaRemove",
		config.Rotas["PessoaRemove"].Tipo,
		config.Rotas["PessoaRemove"].Rota,
		controller.PessoaRemove,
		true,
	},
	Route{
		"PessoaAlter",
		config.Rotas["PessoaAlter"].Tipo,
		config.Rotas["PessoaAlter"].Rota,
		controller.PessoaAlter,
		true,
	},
	Route{
		"PessoaEstado",
		config.Rotas["PessoaEstado"].Tipo,
		config.Rotas["PessoaEstado"].Rota,
		controller.PessoaEstado,
		true,
	},
	Route{
		"PessoaAdmin",
		config.Rotas["PessoaAdmin"].Tipo,
		config.Rotas["PessoaAdmin"].Rota,
		controller.PessoaAdmin,
		true,
	},
}
