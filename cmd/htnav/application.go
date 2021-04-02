// Package htnav exposes the navstar system via HTTP.
package htnav

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/navstar"
)

func NewApplication(sys *navstar.System) *Application {
	return &Application{sys: sys}
}

type Application struct {
	sys *navstar.System
}

// Router returns a router providing HTTP access to the navstar
// system.
func (me *Application) Router() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/routes", me.serveRoutes)
	return m
}

func (me *Application) serveRoutes(w http.ResponseWriter, r *http.Request) {
	// Default to the passenger role
	var user navstar.User
	role := user.Use(me.sys, getRole(r))

	routes, _ := role.ListFlightplans()
	json.NewEncoder(w).Encode(routes)
}

func getRole(r *http.Request) navstar.Role {
	switch r.URL.Query().Get("role") {
	case "pilot":
		return &navstar.Pilot{}
	default:
		return &navstar.Passenger{}
	}
}
