package models

// Route - List all the stops associated to the same line
type Route struct {
    Line string `json:"line"`
    Stops []Transport `json:"stops"`
    Color string `json:"color"`
}
