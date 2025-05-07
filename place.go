package pana

import (
	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
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
