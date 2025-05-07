package pana

import (
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Image is the ActivityStreams Image type.
//
// It shares all properties with [Document].
type Image = Document

func NewImage() *Image {
	return (*Image)(NewDocument().SetType(as.TypeImage))
}
