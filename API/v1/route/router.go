package route

import (
	"controle_pessoal_de_financas/API/v1/controller"
	"controle_pessoal_de_financas/API/v1/logger"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return controller.GetMySigningKey(), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range MyRoutes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.ServeHttpAndLog(handler, route.Name)
		if route.Auth {
			handler = jwtMiddleware.Handler(handler)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
