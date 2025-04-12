package pana

import as "code.dny.dev/pana/vocab/w3/activitystreams"

// Mention is the ActivityStreams Mention.
type Mention = LinkTag

// NewMention initialises a new Mention.
func NewMention() *Mention {
	return (*Mention)(NewLinkTag().SetType(as.TypeMention))
}
