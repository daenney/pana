package pana

import (
	"encoding/json/jsontext"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Source is the ActivityStreams Source type.
//
// It holds the source from which the content of an object was derived.
//
// See https://www.w3.org/TR/activitypub/#source-property.
type Source ld.Node

// NewSource initialises a new Source.
func NewSource() *Source {
	return &Source{
		Properties: make(ld.Properties, 2),
	}
}

// Build finalises the Source.
func (s *Source) Build() Source {
	return *s
}

// See [Object.GetContent].
func (s *Source) GetContent() iter.Seq[*Localised] {
	return (*Object)(s).GetContent()
}

// See [Object.AddContent].
func (s *Source) AddContent(ls ...Localised) *Source {
	(*Object)(s).AddContent(ls...)
	return s
}

// GetMediaType returns the value in [as.MediaType].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-mediatype.
func (s *Source) GetMediaType() jsontext.Value {
	if nodes := (*ld.Node)(s).GetNodes(as.MediaType); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetMediaType sets the string in [as.MediaType].
func (s *Source) SetMediaType(v string) *Source {
	data, _ := jsontext.AppendQuote(nil, v)
	(*ld.Node)(s).SetNodes(as.MediaType, ld.Node{Value: data})
	return s
}

// SetMediaTypeRaw sets the value in [as.MediaType].
func (s *Source) SetMediaTypeRaw(v jsontext.Value) *Source {
	(*ld.Node)(s).SetNodes(as.MediaType, ld.Node{Value: v})
	return s
}
