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

	// TransportType - Structure containing information for a transport type
	TransportType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon,omitempty"`
	}

	// Information - structure for transports informations (passage, nb available)
	Information struct {
		Title     string   `json:"title"`               // Direction of the passage
		Content   []string `json:"content"`             // Time, is array of string to support non numeric values
		Warn      bool     `json:"warn,omitempty"`      // Display the content differently to warn the user
		Timestamp int      `json:"timestamp,omitempty"` // The time at wich the content was updated
	}

	// Transport - Can be embedded by custom Transports structs
	// Gives some usefull properties and two methods
	// Each Transports struct still need to implement UpdateInfo
	Transport struct {
		ID           string        `json:"id"`             // ID of the Transport, should be specific to the Agency
		AgencyID     string        `json:"agencyID"`       // ID of the associated agency
		Type         int           `json:"type"`           // String identifing the kind of transport
		Name         string        `json:"name"`           // The name of the transport, doesn't have to be unique
		Line         string        `json:"line,omitempty"` // The line of the transport
		Position     Position      `json:"position"`       // Position of the transport
		Informations []Information `json:"informations"`   // Informations on the transports
		IconURL      string        `json:"iconURL"`        // URL to the icon
	}
)

// DistanceFrom - Compute the distance between the transport point and a position
// @param p: the position
// @return the distance
func (t *Transport) DistanceFrom(position Position) float64 {
	return t.Position.DistanceFrom(position)
}
