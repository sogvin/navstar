package spaceflight

type Role interface {
	setUser(v *user)
}

type Pilot struct{ *user }

func (me *Pilot) setUser(v *user) { me.user = v }

type Passenger struct{ *user }

func (me *Passenger) setUser(v *user) { me.user = v }

// ----------------------------------------

// user implements domain logic for operations on the System and
// should be used through a role.
type user struct {
	sys *System
}

func (me *user) setSystem(v *System) { me.sys = v }

// ----------------------------------------

// PlanRoute stores the given route. Fails if route already exists.
func (me *Pilot) PlanRoute(v Route) error { return me.planRoute(v) }

func (me *user) planRoute(route Route) error {
	// implement...
	return nil
}

// ----------------------------------------

// ListRoutes returns all routes in the system.
func (me *Pilot) ListRoutes() []Route     { return me.listRoutes() }
func (me *Passenger) ListRoutes() []Route { return me.listRoutes() }

func (me *user) listRoutes() []Route {
	// implement list routes
	return nil
}
