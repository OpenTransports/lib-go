package models

import "fmt"

// Agency - Can be embedded by custom Agencies structs
// The gives some properties and to methods
type Agency struct {
	ID         string           `json:"id"`     // ID of the region (Country.City.Agency)
	Name       string           `json:"name"`   // Displayed name of the Agency
	URL        string           `json:"url"`    // The URL to the agency's website/app...
	Center     Position         `json:"center"` // Center of the Agency
	Radius     float64          `json:"radius"` // Radius of the Agency in meters
	Types      map[int]TypeInfo `json:"types"`  // The type of transports handled by the agency
	Transports []Transport      `json:"-"`      // List of all Transports for the agency
	Routes     []Route          `json:"-"`      // List of all Routes for the agency
}

// String - Stringify an agency
// @return the agency's ID
func (a Agency) String() string {
	return fmt.Sprintf("%v: %v o--------> %v m", a.Name, a.Center, a.Radius)
}

// ContainsPosition - Tell if a position is in the bounds of the agency
// @return the boolean, true if the position is in the agency's bounds
// TODO - Handle the case where the Agency is on the latitude or longitude 0
func (a Agency) ContainsPosition(position Position) bool {
	return a.Center.DistanceFrom(position) <= a.Radius
}

// TransportsNearPosition - Find transports contained in a circle
// @param p: the center of the circle
// @param radius: the radius of the circle in meters
// @return the transports list
func (a Agency) TransportsNearPosition(position Position, radius int) []Transport {
	// Init array of filtered transports
	filteredTransports := make([]Transport, 0, 200)
	// Loop trough agencies transports to find the one that are in the radius limits
	i := 0
	for j, t := range a.Transports {
		if t.DistanceFrom(position) < float64(radius) {
			filteredTransports = append(filteredTransports, a.Transports[j])
			i++
		}
	}
	// Return the transport slice from 0 to i
	// Indexes after i might be empty spots in the original array
	return filteredTransports[:i]
}

// RouteForTransports -
func (a Agency) RouteForTransports(transport Transport) *Route {
	for _, route := range a.Routes {
		if route.Line == transport.Line {
			return &route
		}
	}

	return nil
}

// TransportForID -
func (a Agency) TransportForID(ID string) *Transport {
	for _, transport := range a.Transports {
		if transport.ID == ID {
			return &transport
		}
	}

	return nil
}
