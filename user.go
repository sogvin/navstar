package navstar

// User implements domain logic for operations on the System and
// should be used through a role.
type User struct {
	sys *System
}

func (me *User) Use(v *System, role Role) {
	me.sys = v
	role.setUser(me)
}

// Keep all feature methods private and expose only through roles

func (me *User) submitFlightplan(route Flightplan) error {
	// implement...
	return nil
}

func (me *User) listFlightplans() ([]Flightplan, error) {
	// implement list routes
	return nil, nil
}

// ----------------------------------------

// User implements Role interface with all feature methods returning
// ErrUnauthorized. Role implementations that embed the user only need
// to implement feature methods that are allowed.

func (me *User) SubmitFlightplan(v Flightplan) error {
	return ErrUnauthorized
}

func (me *User) ListFlightplans() ([]Flightplan, error) {
	return nil, ErrUnauthorized
}
