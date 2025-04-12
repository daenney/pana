package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Page is the ActivityStreams Page type.
type Page Object

// NewPage initialises a new Page.
func NewPage() *Page {
	return &Page{
		Properties: make(ld.Properties),
		Type:       []string{as.TypePage},
	}
}

// Build finalises the Page.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (p *Page) Build() Any {
	return Any(*p)
}
