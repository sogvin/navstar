package navstar

type Flightplan struct {
	Routes []Route
}

type Route struct {
	Waypoints []Waypoint
}

type Waypoint struct{}
