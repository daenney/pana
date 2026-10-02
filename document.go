package pana

import (
	"encoding/json/jsontext"
	"iter"
	"strconv"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/mastodon"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
)

// Document is the ActivityStreams Document type.
//
// It shares all properties with [Object].
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
func (d *Document) GetBlurhash() jsontext.Value {
	if nodes := (*ld.Node)(d).GetNodes(mastodon.Blurhash); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetBlurhash sets the value in [mastodon.Blurhash].
func (d *Document) SetBlurhash(v jsontext.Value) *Document {
	(*ld.Node)(d).SetNodes(mastodon.Blurhash, ld.Node{Value: v})
	return d
}

// GetDuration returns the value in [as.Duration].
func (d *Document) GetDuration() jsontext.Value {
	if nodes := (*ld.Node)(d).GetNodes(as.Duration); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetDuration sets the value in [as.Duration].
//
// This value is the XML duration format, which cannot currently be serialised
// from a [time.Duration]. See https://github.com/golang/go/issues/71631.
func (d *Document) SetDuration(v jsontext.Value) *Document {
	(*ld.Node)(d).SetNodes(as.Duration, ld.Node{Value: v, Type: []string{xmlschema.TypeDuration}})
	return d
}

// GetFocalPoint returns the value in [mastodon.FocalPoint].
func (d *Document) GetFocalPoint() []jsontext.Value {
	if nodes := (*ld.Node)(d).GetNodes(mastodon.FocalPoint); len(nodes) == 1 && len(nodes[0].List) == 2 {
		x := nodes[0].List[0].Value
		y := nodes[0].List[1].Value
		return []jsontext.Value{x, y}
	}

	return nil
}

// SetFocalPoint sets the X and Y coordinates in [mastodon.FocalPoint].
func (d *Document) SetFocalPoint(x, y float32) *Document {
	dx := jsontext.AppendFloat(nil, float64(x), 32)
	dy := jsontext.AppendFloat(nil, float64(y), 32)
	(*ld.Node)(d).SetNodes(mastodon.FocalPoint, ld.Node{
		List: []ld.Node{{Value: dx}, {Value: dy}},
	})
	return d
}

// SetFocalPointRaw sets the value in [mastodon.FocalPoint].
func (d *Document) SetFocalPointRaw(x, y jsontext.Value) *Document {
	(*ld.Node)(d).SetNodes(mastodon.FocalPoint, ld.Node{
		List: []ld.Node{{Value: x}, {Value: y}},
	})
	return d
}

// GetHeight returns the value in [as.Height].
func (d *Document) GetHeight() jsontext.Value {
	if nodes := (*ld.Node)(d).GetNodes(as.Height); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetHeight sets the height in [as.Height].
func (d *Document) SetHeight(v uint64) *Document {
	data := strconv.AppendUint(nil, v, 10)
	(*ld.Node)(d).SetNodes(as.Height, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// SetHeightRaw sets the value in [as.Height].
func (d *Document) SetHeightRaw(v jsontext.Value) *Document {
	(*ld.Node)(d).SetNodes(as.Height, ld.Node{Value: v, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// See [Object.GetMediaType].
func (d *Document) GetMediaType() jsontext.Value {
	return (*Object)(d).GetMediaType()
}

// See [Object.SetMediaType].
func (d *Document) SetMediaType(v string) *Document {
	(*Object)(d).SetMediaType(v)
	return d
}

// See [Object.SetMediaTypeRaw].
func (d *Document) SetMediaTypeRaw(v jsontext.Value) *Document {
	(*Object)(d).SetMediaTypeRaw(v)
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

// See [Object.GetSensitive].
func (d *Document) GetSensitive() jsontext.Value {
	return (*Object)(d).GetSensitive()
}

// See [Object.SetSensitive].
func (d *Document) SetSensitive(v bool) *Document {
	(*Object)(d).SetSensitive(v)
	return d
}

// See [Object.SetSensitiveRaw].
func (d *Document) SetSensitiveRaw(v jsontext.Value) *Document {
	(*Object)(d).SetSensitiveRaw(v)
	return d
}

// GetWidth returns the value in [as.Width].
func (d *Document) GetWidth() jsontext.Value {
	if nodes := (*ld.Node)(d).GetNodes(as.Width); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetWidth sets the width in [as.Width].
func (d *Document) SetWidth(v uint64) *Document {
	data := strconv.AppendUint(nil, v, 10)
	(*ld.Node)(d).SetNodes(as.Width, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return d
}

// SetWidthRaw sets the value in [as.Width].
func (d *Document) SetWidthRaw(v jsontext.Value) *Document {
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
