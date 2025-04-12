package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
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
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (e *Event) Build() Any {
	return Any(*e)
}
