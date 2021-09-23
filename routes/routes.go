package routes

import (
	"net/http"

	"github.com/isgo-golgo13/go-gorilla-restsvc-postgres/route_handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

/** Add Routes **/
var Route_Entries = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/health-check",
		route_handlers.HealthCheck,
	},
	Route{
		"Engines",
		"GET",
		"/engines",
		route_handlers.GetEngines,
	},
	Route{
		"Engine",
		"GET",
		"/engines/{id}",
		route_handlers.GetEngine,
	},
}