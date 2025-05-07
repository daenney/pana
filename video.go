package pana

import (
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Video is the ActivityStreams Video type.
//
// It is the same as [Audio].
type Video = Audio

// NewVideo initialises a new Video.
func NewVideo() *Video {
	return (*Video)(NewAudio().SetType(as.TypeVideo))
}
