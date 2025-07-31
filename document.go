package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/mastodon"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
)

// Document is the ActivityStreams Document type.
//
// It shares all properties with [Link].
type Document Object

// NewDocument initialises a new Document.
func NewDocument() *Document {
	return &Document{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeDocument},
	}
}

// Build finalises the Document.
func (d *Document) Build() Document {
	return *d
}

// See [Object.GetType].
func (d *Document) GetType() string {
	return (*Object)(d).GetType()
}

// See [Object.SetType].
func (d *Document) SetType(typ string) *Document {
	(*Object)(d).SetType(typ)
	return d
}

// GetBlurhash returns the value in [mastodon.Blurhash].
func (d *Document) GetBlurhash() json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(mastodon.Blurhash); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetBlurhash sets the value in [mastodon.Blurhash].
func (d *Document) SetBlurhash(v json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(mastodon.Blurhash, ld.Node{Value: v})
	return d
}

// GetDuration returns the value in [as.Duration].
func (d *Document) GetDuration() json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(as.Duration); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetDuration sets the value in [as.Duration].
//
// This value is the XML duration format, which cannot currently be serialised
// from a [time.Duration]. See https://github.com/golang/go/issues/71631.
func (d *Document) SetDuration(v json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(as.Duration, ld.Node{Value: v, Type: []string{xmlschema.TypeDuration}})
	return d
}

// GetFocalPoint returns the value in [mastodon.FocalPoint].
func (d *Document) GetFocalPoint() []json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(mastodon.FocalPoint); len(nodes) == 1 && len(nodes[0].List) == 2 {
		x := nodes[0].List[0].Value
		y := nodes[0].List[1].Value
		return []json.RawMessage{x, y}
	}

	return nil
}

// SetFocalPoint sets the X and Y coordinates in [mastodon.FocalPoint].
func (d *Document) SetFocalPoint(x, y float32) *Document {
	dx, _ := json.Marshal(x)
	dy, _ := json.Marshal(y)
	(*ld.Node)(d).SetNodes(mastodon.FocalPoint, ld.Node{
		List: []ld.Node{{Value: dx}, {Value: dy}},
	})
	return d
}

// SetFocalPointRaw sets the value in [mastodon.FocalPoint].
func (d *Document) SetFocalPointRaw(x, y json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(mastodon.FocalPoint, ld.Node{
		List: []ld.Node{{Value: x}, {Value: y}},
	})
	return d
}

// GetHeight returns the value in [as.Height].
func (d *Document) GetHeight() json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(as.Height); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetHeight sets the height in [as.Height].
func (d *Document) SetHeight(v uint64) *Document {
	data, _ := json.Marshal(v)
	(*ld.Node)(d).SetNodes(as.Height, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// SetHeightRaw sets the value in [as.Height].
func (d *Document) SetHeightRaw(v json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(as.Height, ld.Node{Value: v, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// GetMediaType returns the value in [as.MediaType].
func (d *Document) GetMediaType() json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(as.MediaType); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetMediaType sets the string in [as.MediaType].
func (d *Document) SetMediaType(v string) *Document {
	data, _ := json.Marshal(v)
	(*ld.Node)(d).SetNodes(as.MediaType, ld.Node{Value: data})
	return d
}

// SetMediaTypeRaw sets the value in [as.MediaType].
func (d *Document) SetMediaTypeRaw(v json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(as.MediaType, ld.Node{Value: v})
	return d
}

// See [Object.GetName].
func (d *Document) GetName() iter.Seq[*Localised] {
	return (*Object)(d).GetName()
}

// See [Object.AddName].
func (d *Document) AddName(ls ...Localised) *Document {
	(*Object)(d).AddName(ls...)
	return d
}

// GetWidth returns the value in [as.Width].
func (d *Document) GetWidth() json.RawMessage {
	if nodes := (*ld.Node)(d).GetNodes(as.Width); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetWidth sets the width in [as.Width].
func (d *Document) SetWidth(v uint64) *Document {
	data, _ := json.Marshal(v)
	(*ld.Node)(d).SetNodes(as.Width, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// SetWidthRaw sets the value in [as.Width].
func (d *Document) SetWidthRaw(v json.RawMessage) *Document {
	(*ld.Node)(d).SetNodes(as.Width, ld.Node{Value: v, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// See [Object.GetURL].
func (d *Document) GetURL() string {
	return (*Object)(d).GetURL()
}

// See [Object.SetURL].
func (d *Document) SetURL(url string) *Document {
	(*Object)(d).SetURL(url)
	return d
}
