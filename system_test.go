package navstar

import "testing"

func Test_Pilot(t *testing.T) {
	var role Pilot
	var user User
	user.Use(NewSystem(), &role)

	if _, err := role.ListFlightplans(); err != nil {
		t.Fatal(err)
	}
	if err := role.SubmitFlightplan(Flightplan{}); err != nil {
		t.Fatal(err)
	}
}

func Test_Passenger(t *testing.T) {
	var role Passenger
	var user User
	user.Use(NewSystem(), &role)

	if _, err := role.ListFlightplans(); err != nil {
		t.Fatal(err)
	}
	if err := role.SubmitFlightplan(Flightplan{}); err == nil {
		t.Fatal("expected unauthorized")
	}
}
