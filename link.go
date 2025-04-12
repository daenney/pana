package pana

import (
	"encoding/json"
	"iter"

	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Link is the ActivityStreams Link type.
//
// Though Link can have an ID, this never happens in practice because the target
// of a link is stored in [as.Href] instead.
type Link ld.Node

// NewLink initialises a new Link.
func NewLink() *Link {
	return &Link{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeLink},
	}
}

// Build finalises the Link.
func (l *Link) Build() Link {
	return *l
}

// See [Object.GetType].
//
// This will most commonly be one of:
//   - [as.TypeLink]
//   - [as.TypeDocument]
//   - [as.TypeImage]
//   - [as.TypeAudio]
//   - [as.TypeVideo]
//   - [as.TypeMention]
//   - [as.TypeHashtag]
//
// You'll probably want to use [LinkTag] for Mention and Hashtag.
func (l *Link) GetType() string {
	if l.Type == nil {
		return ""
	}
	return l.Type[0]
}

// See [Object.SetType].
func (l *Link) SetType(t string) *Link {
	l.Type = []string{t}
	return l
}

// GetHref returns the value from [as.Href].
func (l *Link) GetHref() string {
	if nodes := (*ld.Node)(l).GetNodes(as.Href); len(nodes) == 1 {
		return nodes[0].ID
	}
	return ""
}

// SetHref sets a URL in [as.Href].
func (l *Link) SetHref(url string) *Link {
	(*ld.Node)(l).SetNodes(as.Href, ld.Node{ID: url})
	return l
}

// GetName returns the value from [as.Name].
func (l *Link) GetName() iter.Seq[*Localised] {
	return func(yield func(*Localised) bool) {
		for _, n := range (*ld.Node)(l).GetNodes(as.Name) {
			if !yield((*Localised)(&n)) {
				return
			}
		}
	}
}

// SetName sets a value in [as.Name].
func (l *Link) AddName(ls ...Localised) *Link {
	(*ld.Node)(l).AddNodes(as.Name, toLDNodes(ls...)...)
	return l
}

// GetHreflang returns the value from [as.Hreflang].
func (l *Link) GetHreflang() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.Hreflang); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetHreflang sets a value in [as.Hreflang].
func (l *Link) SetHreflang(hreflang json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Hreflang, ld.Node{Value: hreflang})
	return l
}

// GetMediaType returns the value from [as.MediaType].
func (l *Link) GetMediaType() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.MediaType); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetMediaType sets a value in [as.MediaType].
func (l *Link) SetMediaType(mediaType json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.MediaType, ld.Node{Value: mediaType})
	return l
}

// GetRel returns the value from [as.Rel].
func (l *Link) GetRel() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.Rel); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetRel sets a value in [as.Rel].
func (l *Link) SetRel(rel json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Rel, ld.Node{Value: rel})
	return l
}

// GetHeight returns the value from [as.Height].
func (l *Link) GetHeight() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.Height); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetHeight sets a value in [as.Height].
func (l *Link) SetHeight(height json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Height, ld.Node{Value: height})
	return l
}

// GetWidth returns the value from [as.Width].
func (l *Link) GetWidth() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.Width); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetWidth sets a value in [as.Width].
func (l *Link) SetWidth(width json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Width, ld.Node{Value: width})
	return l
}

// LinkTag is a more constrained version of [Link] that is used in [as.Tag].
//
// It's usually used to represent hashtags and mentions.
type LinkTag Link

func NewLinkTag() *LinkTag {
	return &LinkTag{
		Properties: make(ld.Properties, 2),
	}
}

// See [Link.GetType].
func (l *LinkTag) GetType() string {
	return (*Link)(l).GetType()
}

// See [Link.SetType].
func (l *LinkTag) SetType(t string) *LinkTag {
	(*Link)(l).SetType(t)
	return l
}

// See [Link.GetHref]
func (l *LinkTag) GetHref() string {
	return (*Link)(l).GetHref()
}

// See [Link.SetHref].
func (l *LinkTag) SetHref(url string) *LinkTag {
	(*Link)(l).SetHref(url)
	return l
}

// See [Link.GetName].
func (l *LinkTag) GetName() iter.Seq[*Localised] {
	return (*Link)(l).GetName()
}

// See [Link.AddName].
func (l *LinkTag) AddName(ls ...Localised) *LinkTag {
	(*Link)(l).AddName(ls...)
	return l
}
