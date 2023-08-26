package entity

type RouteRepository interface {
	Create(route *Route) error
	List() error
}

type Route struct {
	ID          string
	Name        string
	Source      *Coord
	Destination *Coord
}

type Coord struct {
	Lat float64
	Lng float64
}

func NewRoute(id, name string, source *Coord, destination *Coord) *Route {
	return &Route{
		ID:          id,
		Name:        name,
		Source:      source,
		Destination: destination,
	}
}

func NewCoord(lat float64, lng float64) *Coord {
	return &Coord{
		Lat: lat,
		Lng: lng,
	}
}
