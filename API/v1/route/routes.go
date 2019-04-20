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
	// Route{
	// 	"UsuarioShowAdmin",
	// 	"GET",
	// 	"/usuarios/{adminNome}/{usuarioNome}",
	// 	controller.UsuarioShowAdmin,
	// 	true,
	// },
	// Route{
	// 	"UsuarioCreate",
	// 	"POST",
	// 	"/usuarios",
	// 	controller.UsuarioCreate,
	// 	true,
	// },
	// Route{
	// 	"UsuarioRemove",
	// 	"DELETE",
	// 	"/usuarios/{usuarioNome}",
	// 	controller.UsuarioRemove,
	// 	true,
	// },
	// Route{
	// 	"UsuarioAlter",
	// 	"PUT",
	// 	"/usuarios/{usuarioNome}",
	// 	controller.UsuarioAlter,
	// 	true,
	// },
}
