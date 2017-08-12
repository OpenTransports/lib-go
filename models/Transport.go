package models

// Differents transports types
const (
	Tram = iota
	Metro
	Rail
	Bus
	Ferry
	Cable
	Gondola
	Funicular
	Bike
	Car
	Unknown
)

// Map transports types to string
const (
	TramString      = "Tram"
	MetroString     = "Metro"
	RailString      = "Train"
	BusString       = "Bus"
	FerryString     = "Ferry"
	CableString     = "Cable"
	GondolaString   = "Gondola"
	FunicularString = "Funicular"
	BikeString      = "Bike"
	CarString       = "Car"
	UnknownString   = "Unknown"
)

type (
	// Passage - struct for public transports times
	Passage struct {
		Direction string   `json:"direction"` // Direction of the passage
		Times     []string `json:"times"`     // Time, is array of string to support non numeric values
	}

	// Transport - Can be embedded by custom Transports structs
	// Gives some usefull properties and two methods
	// Each Transports struct still need to implement UpdateInfo
	Transport struct {
		ID        string    `json:"ID"`        // ID of the Transport, should be specific to the Agency
		AgencyID  string    `json:"agencyID"`  // ID of the associated agency
		Type      int       `json:"type"`      // String identifing the kind of transport
		Name      string    `json:"name"`      // The name of the transport, doesn't have to be unique
		Line      string    `json:"line"`      // The line of the transport
		Position  Position  `json:"position"`  // Position of the transport
		Passages  []Passage `json:"passages"`  // Next passage for public transports
		IconURL   string    `json:"iconURL"`   // URL to the icon
		Available int       `json:"available"` // Number of available transports for Bikes or Cars
		Empty     int       `json:"empty"`     // Number of empty spots for Bikes or Cars
	}
)

// DistanceFrom - Compute the distance between the transport point and a position
// @param p: the position
// @return the distance
func (t *Transport) DistanceFrom(position Position) float64 {
	return t.Position.DistanceFrom(position)
}
