package navstar

type Role interface {
	SubmitFlightplan(route Route) error
	ListRoutes() ([]Route, error)
	setUser(v *User)
}

// ----------------------------------------

type Pilot struct {
	*User
}

func (me *Pilot) SubmitFlightplan(v Route) error {
	return me.submitFlightplan(v)
}

func (me *Pilot) ListRoutes() ([]Route, error) {
	return me.listRoutes()
}

func (me *Pilot) setUser(v *User) {
	me.User = v
}

// ----------------------------------------

type Passenger struct {
	*User
}

// SubmitFlightplan always returns ErrUnauthorized
func (me *Passenger) SubmitFlightplan(v Route) error {
	return ErrUnauthorized
}

func (me *Passenger) ListRoutes() ([]Route, error) {
	return me.listRoutes()
}

func (me *Passenger) setUser(v *User) {
	me.User = v
}
