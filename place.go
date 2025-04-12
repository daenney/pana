package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Place is the ActivityStreams Place type.
type Place Object

// NewPlace initialises a new Place.
func NewPlace() *Place {
	return &Place{
		Properties: make(ld.Properties),
		Type:       []string{as.TypePlace},
	}
}

// Build finalises the Place.
func (p *Place) Build() Place {
	return *p
}
