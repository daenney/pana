package pana

import (
	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Profile is the ActivityStreams Profile type.
type Profile Object

// NewProfile initialises a new Profile.
func NewProfile() *Profile {
	return &Profile{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeProfile},
	}
}

// Build finalises the Profile.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (p *Profile) Build() Any {
	return Any(*p)
}
