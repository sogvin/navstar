package navstar

// User implements domain logic for operations on the System and
// should be used through a role.
type User struct {
	sys *System
}

func (me *User) Use(v *System, role Role) Role {
	me.sys = v
	role.setUser(me)
	return role
}

// Keep all feature methods private and expose only through roles

func (me *User) submitFlightplan(route Route) error {
	// implement...
	return nil
}

func (me *User) listFlightplans() ([]Route, error) {
	// implement list routes
	return nil, nil
}
