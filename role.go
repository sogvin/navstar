package navstar

type Role interface {
	setUser(v *User)
	PlanRoute(route Route) error
	ListRoutes() ([]Route, error)
}

type Pilot struct{ *User }

func (me *Pilot) setUser(v *User) { me.User = v }

func (me *Pilot) PlanRoute(v Route) error      { return me.planRoute(v) }
func (me *Pilot) ListRoutes() ([]Route, error) { return me.listRoutes() }

// ----------------------------------------

type Passenger struct{ *User }

func (me *Passenger) setUser(v *User) { me.User = v }

func (me *Passenger) PlanRoute(v Route) error      { return ErrUnauthorized }
func (me *Passenger) ListRoutes() ([]Route, error) { return me.listRoutes() }

// ----------------------------------------

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

func (me *User) planRoute(route Route) error {
	// implement...
	return nil
}

func (me *User) listRoutes() ([]Route, error) {
	// implement list routes
	return nil, nil
}
