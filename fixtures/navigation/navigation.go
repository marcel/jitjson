package navigation

import "time"

type Route struct {
	Summary string `json:"summary"`
	Legs    []Leg  `json:"legs"`
}

type Leg struct {
	Distance      int           `json:"distance"`
	Duration      time.Duration `json:"duration"`
	StartAddress  Address       `json:"start_address"`
	EndAddress    Address       `json:"end_address"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
}

type TravelMode string

const (
	DrivingMode TravelMode = "DRIVING"
	WalkingMode TravelMode = "WALKING"
)

type Maneuver string

const (
	TurnLeftManeuver  Maneuver = "turn-left"
	TurnRightManeuver Maneuver = "turn-right"
	RampLeftManeuver  Maneuver = "ramp-left"
	RampRightManeuver Maneuver = "ramp-right"
)

type Step struct {
	Distance      int           `json:"distance"`
	Duration      time.Duration `json:"duration"`
	StartLocation Location      `json:"start_location"`
	EndLocation   Location      `json:"end_location"`
	TravelMode    `json:"travel_mode"`
	Maneuver      `json:"maneuver"`
	Instructions  string `json:"instructions"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Address struct {
	Number string
	Street string
}
