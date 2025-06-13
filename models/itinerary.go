package models

import "time"

type Itinerary struct {
	ID          int
	Name        string
	Description string
	Destination string
	DateTime    time.Time
	UserID      int
}

var itineraries = []Itinerary{}

func (e Itinerary) Save() {
	//later add to database
	itineraries = append(itineraries, e)
}

func GetAllItineraries() []Itinerary {
	return itineraries
}
