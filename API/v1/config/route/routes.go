package route

import (
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

var routes = Routes{
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/",
	// 	controller.Index,
	// 	false,
	// },
	// Route{
	// 	"GetStaticDist",
	// 	"GET",
	// 	"/dist/{nomeStatic}",
	// 	controller.GetStatic,
	// 	false,
	// },
	Route{
		"Login",
		"POST",
		"/login/{usuario}",
		controller.Login,
		false,
	},
	Route{
		"TokenValido",
		"GET",
		"/token",
		controller.TokenValido,
		true,
	},
	Route{
		"PessoaIndex",
		"GET",
		"/pessoas",
		controller.PessoaIndex,
		true,
	},
	// Route{
	// 	"UsuarioShow",
	// 	"GET",
	// 	"/usuarios/{usuarioNome}",
	// 	controller.UsuarioShow,
	// 	true,
	// },
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
