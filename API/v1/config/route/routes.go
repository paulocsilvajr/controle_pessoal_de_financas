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
	// Route{
	// 	"AbastecimentoIndex",
	// 	"GET",
	// 	"/abastecimentos",
	// 	controller.AbastecimentoIndex,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoUsuarioMes",
	// 	"GET",
	// 	"/abastecimentos/mes/{mes}-{ano}",
	// 	controller.AbastecimentoUsuarioMes,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoUsuarioMesPorUsuario",
	// 	"GET",
	// 	"/abastecimentos/{usuarioNome}/mes/{mes}-{ano}",
	// 	controller.AbastecimentoUsuarioMesPorUsuario,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoShow",
	// 	"GET",
	// 	"/abastecimentos/{abastecimentoId}",
	// 	controller.AbastecimentoShow,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoCreate",
	// 	"POST",
	// 	"/abastecimentos",
	// 	controller.AbastecimentoCreate,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoRemove",
	// 	"DELETE",
	// 	"/abastecimentos/{abastecimentoId}",
	// 	controller.AbastecimentoRemove,
	// 	true,
	// },
	// Route{
	// 	"AbastecimentoAlter",
	// 	"PUT",
	// 	"/abastecimentos/{abastecimentoId}",
	// 	controller.AbastecimentoAlter,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemIndex",
	// 	"GET",
	// 	"/quilometragens",
	// 	controller.QuilometragemIndex,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemUsuarioMes",
	// 	"GET",
	// 	"/quilometragens/mes/{mes}-{ano}",
	// 	controller.QuilometragemUsuarioMes,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemUsuarioMesPorUsuario",
	// 	"GET",
	// 	"/quilometragens/{usuarioNome}/mes/{mes}-{ano}",
	// 	controller.QuilometragemUsuarioMesPorUsuario,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemShow",
	// 	"GET",
	// 	"/quilometragens/{quilometragemId}",
	// 	controller.QuilometragemShow,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemCreate",
	// 	"POST",
	// 	"/quilometragens",
	// 	controller.QuilometragemCreate,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemRemove",
	// 	"DELETE",
	// 	"/quilometragens/{quilometragemId}",
	// 	controller.QuilometragemRemove,
	// 	true,
	// },
	// Route{
	// 	"QuilometragemAlter",
	// 	"PUT",
	// 	"/quilometragens/{quilometragemId}",
	// 	controller.QuilometragemAlter,
	// 	true,
	// },
	// Route{
	// 	"ColaboradorIndex",
	// 	"GET",
	// 	"/colaboradores",
	// 	controller.ColaboradorIndex,
	// 	true,
	// },
	// Route{
	// 	"ColaboradorShow",
	// 	"GET",
	// 	"/colaboradores/{colaboradorNome}",
	// 	controller.ColaboradorShow,
	// 	true,
	// },
	// Route{
	// 	"ColaboradorCreate",
	// 	"POST",
	// 	"/colaboradores",
	// 	controller.ColaboradorCreate,
	// 	true,
	// },
	// Route{
	// 	"ColaboradorRemove",
	// 	"DELETE",
	// 	"/colaboradores/{colaboradorNome}",
	// 	controller.ColaboradorRemove,
	// 	true,
	// },
	// Route{
	// 	"ColaboradorAlter",
	// 	"PUT",
	// 	"/colaboradores/{colaboradorNome}",
	// 	controller.ColaboradorAlter,
	// 	true,
	// },
	// Route{
	// 	"SolicitanteIndex",
	// 	"GET",
	// 	"/solicitantes",
	// 	controller.SolicitanteIndex,
	// 	true,
	// },
	// Route{
	// 	"SolicitanteShow",
	// 	"GET",
	// 	"/solicitantes/{solicitanteNome}",
	// 	controller.SolicitanteShow,
	// 	true,
	// },
	// Route{
	// 	"SolicitanteCreate",
	// 	"POST",
	// 	"/solicitantes",
	// 	controller.SolicitanteCreate,
	// 	true,
	// },
	// Route{
	// 	"SolicitanteRemove",
	// 	"DELETE",
	// 	"/solicitantes/{solicitanteNome}",
	// 	controller.SolicitanteRemove,
	// 	true,
	// },
	// Route{
	// 	"SolicitanteAlter",
	// 	"PUT",
	// 	"/solicitantes/{solicitanteNome}",
	// 	controller.SolicitanteAlter,
	// 	true,
	// },
}
