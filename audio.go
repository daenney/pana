package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Audio is the ActivityStreams Audio type.
//
// It is a [Document] without width/height.
type Audio Document

// NewAudio initialises a new Audio.
func NewAudio() *Audio {
	return &Audio{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeAudio},
	}
}

// Build finalises the Audio.
func (a *Audio) Build() Audio {
	return *a
}

// See [Object.GetType].
func (a *Audio) GetType() string {
	return (*Object)(a).GetType()
}

// See [Object.SetType].
func (a *Audio) SetType(typ string) *Audio {
	(*Object)(a).SetType(typ)
	return a
}

// GetDuration returns the value in [as.Duration].
func (a *Audio) GetDuration() json.RawMessage {
	return (*Document)(a).GetDuration()
}

// SetDuration sets the value in [as.Duration].
//
// See [Document.SetDuration] for the format.
func (a *Audio) SetDuration(v json.RawMessage) *Audio {
	(*Document)(a).SetDuration(v)
	return a
}

// GetMediaType returns the value in [as.MediaType].
func (a *Audio) GetMediaType() json.RawMessage {
	return (*Document)(a).GetMediaType()
}

// SetMediaType sets the string in [as.MediaType].
func (a *Audio) SetMediaType(v string) *Audio {
	(*Document)(a).SetMediaType(v)
	return a
}

// SetMediaTypeRaw sets the value in [as.MediaType].
func (a *Audio) SetMediaTypeRaw(v json.RawMessage) *Audio {
	(*Document)(a).SetMediaTypeRaw(v)
	return a
}

// See [Object.GetName].
func (a *Audio) GetName() iter.Seq[*Localised] {
	return (*Object)(a).GetName()
}

// See [Object.AddName].
func (a *Audio) AddName(ls ...Localised) *Audio {
	(*Object)(a).AddName(ls...)
	return a
}

// See [Object.GetURL].
func (a *Audio) GetURL() string {
	return (*Object)(a).GetURL()
}

// See [Object.SetURL].
func (a *Audio) SetURL(url string) *Audio {
	(*Object)(a).SetURL(url)
	return a
}
