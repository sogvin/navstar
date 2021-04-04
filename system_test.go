package navstar

import "testing"

func Test_ListFlightplans(t *testing.T) {
	okRoles := []Role{&Pilot{}, &Passenger{}, &Crew{}}

	for _, role := range okRoles {
		var user User
		user.Use(NewSystem(), role)
		if _, err := role.ListFlightplans(); err != nil {
			t.Error(err)
		}
	}
}

// ----------------------------------------

func Test_SubmitFlightplan(t *testing.T) {
	okRoles := []Role{&Pilot{}}

	for _, role := range okRoles {
		var user User
		user.Use(NewSystem(), role)
		var plan Flightplan
		if err := role.SubmitFlightplan(plan); err != nil {
			t.Error(err)
		}
	}
}

func Test_SubmitFlightplan_fails(t *testing.T) {
	okRoles := []Role{&Passenger{}}

	for _, role := range okRoles {
		var user User
		user.Use(NewSystem(), role)
		var plan Flightplan
		if err := role.SubmitFlightplan(plan); err == nil {
			t.Error("SubmitFlightplan", role, err)
		}
	}
}
