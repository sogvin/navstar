package htapi

import (
	"net/http/httptest"
	"testing"

	"github.com/gregoryv/navstar"
)

func Test_routes(t *testing.T) {
	sys := navstar.NewSystem()
	router := NewRouter(sys)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/flightplans", nil)

	router.ServeHTTP(w, r)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}

func Test_getRole(t *testing.T) {
	role := getRole(httptest.NewRequest("GET", "/flightplans", nil))
	if _, ok := role.(*navstar.Passenger); !ok {
		t.Errorf("not passenger %#v", role)
	}

	role = getRole(httptest.NewRequest("GET", "/flightplans?role=pilot", nil))
	if _, ok := role.(*navstar.Pilot); !ok {
		t.Errorf("not pilot %#v", role)
	}

	role = getRole(httptest.NewRequest("GET", "/flightplans?role=crew", nil))
	if _, ok := role.(*navstar.Crew); !ok {
		t.Errorf("not crew %#v", role)
	}
}
