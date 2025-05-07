package pana

import (
	"iter"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Instrument is the ActivityStreams instrument object.
type Instrument Object

// NewInstrument initialises a new Instrument.
func NewInstrument() *Instrument {
	return &Instrument{
		Properties: make(ld.Properties, 2),
		Type:       []string{as.TypeService},
	}
}

// Build finalises the Instrument.
func (i *Instrument) Build() Instrument {
	return *i
}

// See [Object.GetType].
func (i *Instrument) GetType() string {
	return (*Object)(i).GetType()
}

// See [Object.SetType].
func (i *Instrument) SetType(typ string) *Instrument {
	(*Object)(i).SetType(typ)
	return i
}

// See [Object.GetName].
func (i *Instrument) GetName() iter.Seq[*Localised] {
	return (*Object)(i).GetName()
}

// See [Object.AddName].
func (i *Instrument) AddName(ls ...Localised) *Instrument {
	(*Object)(i).AddName(ls...)
	return i
}

// See [Object.GetURL].
func (i *Instrument) GetURL() string {
	return (*Object)(i).GetURL()
}

// See [Object.SetURL].
func (i *Instrument) SetURL(url string) *Instrument {
	(*Object)(i).SetURL(url)
	return i
}
