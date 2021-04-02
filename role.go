package navstar

type Role interface {
	SubmitFlightplan(route Flightplan) error
	ListFlightplans() ([]Flightplan, error)
	setUser(v *User)
}

// ----------------------------------------

type Pilot struct {
	*User
}

func (me *Pilot) SubmitFlightplan(v Flightplan) error {
	return me.submitFlightplan(v)
}

func (me *Pilot) ListFlightplans() ([]Flightplan, error) {
	return me.listFlightplans()
}

func (me *Pilot) setUser(v *User) {
	me.User = v
}

// ----------------------------------------

type Crew struct {
	*User
}

// SubmitFlightplan always returns ErrUnauthorized
func (me *Crew) SubmitFlightplan(v Flightplan) error {
	return ErrUnauthorized
}

func (me *Crew) ListFlightplans() ([]Flightplan, error) {
	return me.listFlightplans()
}

func (me *Crew) setUser(v *User) {
	me.User = v
}

// ----------------------------------------

type Passenger struct {
	*User
}

// SubmitFlightplan always returns ErrUnauthorized
func (me *Passenger) SubmitFlightplan(v Flightplan) error {
	return ErrUnauthorized
}

func (me *Passenger) ListFlightplans() ([]Flightplan, error) {
	return me.listFlightplans()
}

func (me *Passenger) setUser(v *User) {
	me.User = v
}
