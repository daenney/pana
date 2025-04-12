package pana

import (
	"encoding/json"

	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Icon is a more constrained version of [Image] used with [as.Icon].
type Icon Image

// NewIcon initialises a new Icon.
func NewIcon() *Icon {
	return &Icon{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeImage},
	}
}

// Build finalises the Icon.
func (i *Icon) Build() Icon {
	return *i
}

// See [Object.GetType].
func (i *Icon) GetType() string {
	return (*Image)(i).GetType()
}

// See [Object.SetType].
func (i *Icon) SetType() *Icon {
	(*Image)(i).SetType(as.TypeImage)
	return i
}

// GetMediaType returns the value in [as.MediaType].
func (i *Icon) GetMediaType() json.RawMessage {
	return (*Image)(i).GetMediaType()
}

// SetMediaType sets the value in [as.MediaType].
func (i *Icon) SetMediaType(v json.RawMessage) *Icon {
	(*Image)(i).SetMediaType(v)
	return i
}

// See [Object.GetURL].
func (i *Icon) GetURL() string {
	return (*Image)(i).GetURL()
}

// See [Object.SetURL].
func (i *Icon) SetURL(url string) *Icon {
	(*Image)(i).SetURL(url)
	return i
}
