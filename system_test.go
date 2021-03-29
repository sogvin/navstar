package spaceflight

import "testing"

func Test_Service_Use(t *testing.T) {
	var srv System
	var role Pilot
	srv.Use(&role)
}
