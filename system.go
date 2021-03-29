package spaceflight

// System provides logic for reading and writing spaceflight related
// resources.
type System struct {
	// e.g. database, sync mutexes
}

// Use provides the given role access to the service.
func (me *System) Use(role Role) {
	var u user
	u.setSystem(me)
	role.setUser(&u)
}
