package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

type Instrument Object

// NewInstrument initialises a new Instrument.
func NewInstrument() *Instrument {
	return &Instrument{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeApplication},
	}
}

// Build finalises the Instrument.
func (i *Instrument) Build() Instrument {
	return *i
}
