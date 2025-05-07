package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Audio is the ActivityStreams Audio type.
//
// It is a [Document] without width/height but with duration.
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
	if nodes := (*ld.Node)(a).GetNodes(as.Duration); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetDuration sets the value in [as.Duration].
func (a *Audio) SetDuration(v json.RawMessage) *Audio {
	(*ld.Node)(a).SetNodes(as.Duration, ld.Node{Value: v})
	return a
}

// GetMediaType returns the value in [as.MediaType].
func (a *Audio) GetMediaType() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(as.MediaType); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetMediaType sets the value in [as.MediaType].
func (a *Audio) SetMediaType(v json.RawMessage) *Audio {
	(*ld.Node)(a).SetNodes(as.MediaType, ld.Node{Value: v})
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
