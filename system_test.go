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

func Test_ListFlightplans_fails(t *testing.T) {
	okRoles := []Role{&User{}}

	for _, role := range okRoles {
		var user User
		user.Use(NewSystem(), role)
		if _, err := role.ListFlightplans(); err == nil {
			t.Error("ListFlightplans", role, err)
		}
	}
}

func Test_Pilot(t *testing.T) {
	var role Pilot
	var user User
	user.Use(NewSystem(), &role)

	if err := role.SubmitFlightplan(Flightplan{}); err != nil {
		t.Fatal(err)
	}
}

func Test_Passenger(t *testing.T) {
	var role Passenger
	var user User
	user.Use(NewSystem(), &role)

	expectUnauth(t, role.SubmitFlightplan(Flightplan{}))
}

func Test_Crew(t *testing.T) {
	var role Crew
	var user User
	user.Use(NewSystem(), &role)

	expectUnauth(t, role.SubmitFlightplan(Flightplan{}))
}

func Test_User(t *testing.T) {
	var user User

	_, err := user.ListFlightplans()
	expectUnauth(t, err)
	expectUnauth(t, user.SubmitFlightplan(Flightplan{}))
}

func expectUnauth(t *testing.T, err error) {
	t.Helper()
	if err != ErrUnauthorized {
		t.Error("expected ErrUnauthorized:", err)
	}
}
