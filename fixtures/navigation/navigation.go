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
	Steps         []Step        `json:"steps"`
}

var ExampleRoute = Route{
	Summary: "US-101 S and I-280 S",
	Legs: []Leg{
		Leg{
			Distance: 61718,
			Duration: time.Duration(42 * time.Minute),
			StartAddress: Address{
				Number:  "",
				Street:  "Henry St",
				City:    "San Francisco",
				State:   "CA",
				ZipCode: 94114,
				Country: "USA",
			},
			StartLocation: Location{
				Lat: "37.7681528",
				Lng: "-122.4222248",
			},
			EndAddress: Address{
				Number:  "3500",
				Street:  "Deer Creek Rd",
				City:    "Palo Alto",
				State:   "CA",
				ZipCode: 94304,
				Country: "USA",
			},
			EndLocation: Location{
				Lat: "37.7681528",
				Lng: "-122.4222248",
			},
			Steps: []Step{
				Step{
					Distance: 849,
					Duration: time.Duration(3 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Head South east on 14th St toward Sanchez St",
					TravelMode:   DrivingMode,
					Maneuver:     NoManeuver,
				},
				Step{
					Distance: 192,
					Duration: time.Duration(5 * time.Second),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Turn left onto Valencia St",
					TravelMode:   DrivingMode,
					Maneuver:     TurnLeftManeuver,
				},
				Step{
					Distance: 157,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Turn right onto Duboce Ave",
					TravelMode:   DrivingMode,
					Maneuver:     TurnRightManeuver,
				},
				Step{
					Distance: 247,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Continue onto 13th St",
					TravelMode:   DrivingMode,
					Maneuver:     NoManeuver,
				},
				Step{
					Distance: 53,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Turn right onto Van Ness Ave",
					TravelMode:   DrivingMode,
					Maneuver:     TurnRightManeuver,
				},
				Step{
					Distance: 1025,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Slight right to merge onto US-101 S",
					TravelMode:   DrivingMode,
					Maneuver:     NoManeuver,
				},
				Step{
					Distance: 15527,
					Duration: time.Duration(10 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Keep right to stay on US-101 S, follow signs for San Jose",
					TravelMode:   DrivingMode,
					Maneuver:     KeepRightManeuver,
				},
				Step{
					Distance: 1308,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Take exit for I-380 W toward I-280 San Bruno",
					TravelMode:   DrivingMode,
					Maneuver:     RampRightManeuver,
				},
				Step{
					Distance: 1454,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Continue onto I-380 W",
					TravelMode:   DrivingMode,
					Maneuver:     NoManeuver,
				},
				Step{
					Distance: 980,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Take exit on the left for I-280 S toward San Jose",
					TravelMode:   DrivingMode,
					Maneuver:     RampLeftManeuver,
				},
				Step{
					Distance: 37265,
					Duration: time.Duration(20 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Merge onto I-280 S",
					TravelMode:   DrivingMode,
					Maneuver:     MergeManeuver,
				},
				Step{
					Distance: 633,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Take exit for Page Mill Rd toward Arastradero Rd Palo Alto",
					TravelMode:   DrivingMode,
					Maneuver:     RampRightManeuver,
				},
				Step{
					Distance: 25,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Keep left at the fork, follow signs for Arastradero Rd Veterans Hospital Palo Alto",
					TravelMode:   DrivingMode,
					Maneuver:     ForkLeftManeuver,
				},
				Step{
					Distance: 1190,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Turn left onto County Hwy G3 Page Mill Rd",
					TravelMode:   DrivingMode,
					Maneuver:     TurnLeftManeuver,
				},
				Step{
					Distance: 813,
					Duration: time.Duration(1 * time.Minute),
					StartLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					EndLocation: Location{
						Lat: "37.7681528",
						Lng: "-122.4222248",
					},
					Instructions: "Turn right onto Deer Creek Rd. Destination will be on the right.",
					TravelMode:   DrivingMode,
					Maneuver:     TurnRightManeuver,
				},
			},
		},
	},
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
	KeepRightManeuver Maneuver = "keep-right"
	KeepLeftManeuver  Maneuver = "keep-left"
	MergeManeuver     Maneuver = "merge"
	ForkLeftManeuver  Maneuver = "fork-left"
	ForkRightManeuver Maneuver = "fork-right"

	NoManeuver Maneuver = ""
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
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Number  string `json:"number"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zip_code"`
	Country string `json:"country"`
}

type thisWillBeSkippedBecauseItIsNotExported struct {
	Identifier string `json:"identifier"`
}
