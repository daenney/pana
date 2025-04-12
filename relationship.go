package pana

import ld "code.dny.dev/longdistance"

// Relationship is the ActivityStreams Relationship type.
type Relationship Object

// NewRelationship initialises a new Relationship.
func NewRelationship() *Relationship {
	return &Relationship{
		Properties: make(ld.Properties),
	}
}

// Build finalises the Relationship.
func (r *Relationship) Build() Relationship {
	return *r
}
