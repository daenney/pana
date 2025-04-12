package pana

import (
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Hashtag is the ActivityStreams Hashtag.
type Hashtag = LinkTag

// NewHashtag initialises a new Hashtag.
func NewHashtag() *Hashtag {
	return (*Hashtag)(NewLinkTag().SetType(as.TypeHashtag))
}
