package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/logger"
	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/routes"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.Route_Entries {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}