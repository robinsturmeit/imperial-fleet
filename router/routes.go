package router

import (
	"net/http"

	handler "github.com/robinsturmeit/imperial-fleet/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetSpacecrafts",
		"GET",
		"/spacecrafts",
		handler.GetSpacecrafts,
	},
	Route{
		"GetSpacecraftByProp",
		"GET",
		"/spacecraft/{prop}/{value}",
		handler.GetSpacecraftByProp,
	},
}
