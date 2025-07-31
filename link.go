package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
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

// SetHreflang sets the string in [as.Hreflang].
func (l *Link) SetHreflang(hreflang string) *Link {
	data, _ := json.Marshal(hreflang)
	(*ld.Node)(l).SetNodes(as.Hreflang, ld.Node{Value: data})
	return l
}

// SetHreflangRaw sets a value in [as.Hreflang].
func (l *Link) SetHreflanRaw(hreflang json.RawMessage) *Link {
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

// SetMediaType sets a string in [as.MediaType].
func (l *Link) SetMediaType(mediaType string) *Link {
	data, _ := json.Marshal(mediaType)
	(*ld.Node)(l).SetNodes(as.MediaType, ld.Node{Value: data})
	return l
}

// SetMediaTypeRaw sets a value in [as.MediaType].
func (l *Link) SetMediaTypeRaw(mediaType json.RawMessage) *Link {
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

// SetRel sets a string in [as.Rel].
func (l *Link) SetRel(rel string) *Link {
	data, _ := json.Marshal(rel)
	(*ld.Node)(l).SetNodes(as.Rel, ld.Node{Value: data})
	return l
}

// SetRelRaw sets a value in [as.Rel].
func (l *Link) SetRelRaw(rel json.RawMessage) *Link {
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

// SetHeight sets the height in [as.Height].
func (l *Link) SetHeight(height uint64) *Link {
	data, _ := json.Marshal(height)
	(*ld.Node)(l).SetNodes(as.Height, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return l
}

// SetHeightRaw sets a value in [as.Height].
func (l *Link) SetHeightRaw(height json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Height, ld.Node{Value: height, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return l
}

// GetWidth returns the value from [as.Width].
func (l *Link) GetWidth() json.RawMessage {
	if nodes := (*ld.Node)(l).GetNodes(as.Width); len(nodes) == 1 {
		return nodes[0].Value
	}
	return nil
}

// SetWidth sets the width in [as.Width].
func (l *Link) SetWidth(width uint64) *Link {
	data, _ := json.Marshal(width)
	(*ld.Node)(l).SetNodes(as.Width, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return l
}

// SetWidthRaw sets a value in [as.Width].
func (l *Link) SetWidthRaw(width json.RawMessage) *Link {
	(*ld.Node)(l).SetNodes(as.Width, ld.Node{Value: width, Type: []string{xmlschema.TypeNonNegativeInteger}})
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
