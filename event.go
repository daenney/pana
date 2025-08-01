package pana

import (
	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Event is the Activitystreams Event type.
type Event Object

// NewEvent initialises a new Event.
func NewEvent() *Event {
	return &Event{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeEvent},
	}
}

// Build finalises the Event.
func (e *Event) Build() Event {
	return *e
}
