// Package htapi exposes the navstar system via HTTP.
package htapi

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/navstar"
)

func NewHttpRouter(sys *navstar.System) *HttpRouter {
	mux := http.NewServeMux()
	hr := HttpRouter{sys: sys, mux: mux}
	mux.HandleFunc("/routes", hr.serveRoutes)
	return &hr
}

type HttpRouter struct {
	sys *navstar.System
	mux *http.ServeMux
}

func (me *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	me.mux.ServeHTTP(w, r)
}

func (me *HttpRouter) serveRoutes(w http.ResponseWriter, r *http.Request) {
	var role navstar.Role = getRole(r)
	var user navstar.User
	user.Use(me.sys, role)

	routes, _ := role.ListFlightplans()
	json.NewEncoder(w).Encode(routes)
}

// getRole returns a role implementation based on the incomming
// request. This one just looks for the query parameter role. Defaults
// to the passenger role.
func getRole(r *http.Request) navstar.Role {
	switch r.URL.Query().Get("role") {
	case "pilot":
		return &navstar.Pilot{}
	case "crew":
		return &navstar.Crew{}
	default:
		return &navstar.Passenger{}
	}
}
