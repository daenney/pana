package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Tombstone is the ActivityStreams Tombstone type.
type Tombstone Object

// NewTombstone initialises a new Tombstone.
func NewTombstone() *Tombstone {
	return &Tombstone{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeTombstone},
	}
}

// Build finalises the Tombstone.
func (t *Tombstone) Build() Tombstone {
	return *t
}
