package htspace

import (
	"encoding/json"
	"net/http"

	"github.com/gregoryv/sogvin/example/spaceflight"
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
	var role spaceflight.Passenger
	me.sys.Use(&role)

	routes := role.ListRoutes()
	json.NewEncoder(w).Encode(routes)
}
