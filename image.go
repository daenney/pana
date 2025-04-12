package pana

import (
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Image is the ActivityStreams Image type.
//
// It shares all properties with [Document].
type Image = Document

func NewImage() *Image {
	return (*Image)(NewDocument().SetType(as.TypeImage))
}
