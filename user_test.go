package navstar

import "testing"

func Test_ListFlightplans_fails(t *testing.T) {
	var user User
	if _, err := user.ListFlightplans(); err == nil {
		t.Error("ListFlightplans", err)
	}
}
