package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/mastodon"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

type Emoji Object

// NewEmoji initialises a new Emoji.
func NewEmoji() *Emoji {
	return &Emoji{
		Properties: make(ld.Properties),
		Type:       []string{mastodon.TypeEmoji},
	}
}

// Build finalises the Emoji.
func (e *Emoji) Build() Emoji {
	return *e
}

// See [Object.GetID].
//
// Beware that Emoji's might not have an ID. This is used for instance-local
// emoji.
func (e *Emoji) GetID() string {
	return (*Object)(e).GetID()
}

// See [Object.SetID].
func (e *Emoji) SetID(id string) *Emoji {
	(*Object)(e).SetID(id)
	return e
}

// See [Object.GetType].
func (e *Emoji) GetType() string {
	return (*Object)(e).GetType()
}

// See [Object.SetType].
func (e *Emoji) SetType() *Emoji {
	(*Object)(e).SetType(mastodon.TypeEmoji)
	return e
}

// See [Object.GetName].
func (e *Emoji) GetName() iter.Seq[*Localised] {
	return (*Object)(e).GetName()
}

// See [Object.SetName].
func (e *Emoji) AddName(ls ...Localised) *Emoji {
	(*Object)(e).AddName(ls...)
	return e
}

// See [Object.GetUpdated].
func (e *Emoji) GetUpdated() json.RawMessage {
	return (*Object)(e).GetUpdated()
}

// See [Object.SetUpdated].
func (e *Emoji) SetUpdated(v json.RawMessage) *Emoji {
	(*Object)(e).SetUpdated(v)
	return e
}

// GetIcon returns the [Image] in [as.Icon].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-icon.
func (e *Emoji) GetIcon() *Icon {
	if nodes := (*ld.Node)(e).GetNodes(as.Icon); len(nodes) == 1 {
		return (*Icon)(&nodes[0])
	}

	return nil
}

// SetIcon sets the [Image] in [as.Icon].
func (e *Emoji) SetIcon(img Icon) *Emoji {
	(*ld.Node)(e).SetNodes(as.Icon, ld.Node(img))
	return e
}
