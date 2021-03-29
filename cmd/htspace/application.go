// Package htspace exposes the spaceflight system via HTTP.
package htspace

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/spaceflight"
)

func NewApplication(sys *spaceflight.System) *Application {
	return &Application{sys: sys}
}

type Application struct {
	sys *spaceflight.System
}

// Router returns a router providing HTTP access to the spaceflight
// system.
func (me *Application) Router() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/routes", me.serveRoutes)
	return m
}

func (me *Application) serveRoutes(w http.ResponseWriter, r *http.Request) {
	// Default to the passenger role
	var user spaceflight.User
	role := user.Use(me.sys, getRole(r))

	routes, _ := role.ListRoutes()
	json.NewEncoder(w).Encode(routes)
}

func getRole(r *http.Request) spaceflight.Role {
	switch r.URL.Query().Get("role") {
	case "pilot":
		return &spaceflight.Pilot{}
	default:
		return &spaceflight.Passenger{}
	}
}
