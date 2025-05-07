package pana

import (
	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Tombstone is the ActivityStreams Tombstone type.
type Tombstone Object

// NewTombstone initialises a new Tombstone.
func NewTombstone() *Tombstone {
	return &Tombstone{
		Properties: make(ld.Properties, 1),
		Type:       []string{as.TypeTombstone},
	}
}

// Build finalises the Tombstone.
func (t *Tombstone) Build() Tombstone {
	return *t
}

// See [Object.GetID].
func (t *Tombstone) GetID() string {
	return (*Object)(t).GetID()
}

// See [Object.SetID].
func (t *Tombstone) SetID(id string) *Tombstone {
	(*Object)(t).SetID(id)
	return t
}

// See [Object.GetType].
func (t *Tombstone) GetType() string {
	return (*Object)(t).GetType()
}

// See [Object.SetType].
func (t *Tombstone) SetType(id string) *Tombstone {
	(*Object)(t).SetType(id)
	return t
}

// Returns the type in [as.FormerType].
func (t *Tombstone) GetFormerType() string {
	if nodes := (*ld.Node)(t).GetNodes(as.FormerType); len(nodes) != 0 {
		return nodes[0].ID
	}

	return ""
}

// Sets the type in [as.FormerType].
func (t *Tombstone) SetFormerType(id string) *Tombstone {
	(*ld.Node)(t).SetNodes(as.FormerType, ld.Node{ID: id})
	return t
}
