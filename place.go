package pana

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/schema"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
)

// Place is the ActivityStreams Place type.
type Place Object

// NewPlace initialises a new Place.
func NewPlace() *Place {
	return &Place{
		Properties: make(ld.Properties),
		Type:       []string{as.TypePlace},
	}
}

// Build finalises the Place.
func (p *Place) Build() Place {
	return *p
}

// See [Object.GetID].
func (p *Place) GetID() string {
	return (*Object)(p).GetID()
}

// See [Object.SetID].
func (p *Place) SetID(id string) *Place {
	(*Object)(p).SetID(id)
	return p
}

// See [Object.GetType].
func (p *Place) GetType() string {
	return (*Object)(p).GetType()
}

// See [Object.SetType].
func (p *Place) SetType(typ string) *Place {
	(*Object)(p).SetType(typ)
	return p
}

// GetAddress returns the [PostalAddress] in [schema.Address].
func (p *Place) GetAddress() *PostalAddress {
	if nodes := (*ld.Node)(p).GetNodes(schema.Address); len(nodes) == 1 {
		return (*PostalAddress)(&nodes[0])
	}

	return nil
}

// SetAddress sets the [PostalAddress] in [schema.Address].
func (p *Place) SetAddress(pa PostalAddress) *Place {
	(*ld.Node)(p).SetNodes(schema.Address, ld.Node(pa))
	return p
}

// GetLatitude returns the value in [as.Latitude].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-latitude.
func (p *Place) GetLatitude() jsontext.Value {
	if nodes := (*ld.Node)(p).GetNodes(as.Latitude); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetLatitude sets the value in [as.Latitude].
func (p *Place) SetLatitude(v float32) *Place {
	data, _ := json.Marshal(v)
	(*ld.Node)(p).SetNodes(as.Latitude, ld.Node{Value: data, Type: []string{xmlschema.TypeFloat}})
	return p
}

// SetLatitudeRaw sets the value in [as.Latitude].
func (p *Place) SetLatitudeRaw(v jsontext.Value) *Place {
	(*ld.Node)(p).SetNodes(as.Latitude, ld.Node{Value: v, Type: []string{xmlschema.TypeFloat}})
	return p
}

// GetLongitude returns the value in [as.Longitude].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-longitude.
func (p *Place) GetLongitude() jsontext.Value {
	if nodes := (*ld.Node)(p).GetNodes(as.Longitude); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetLongitude sets the value in [as.Longitude].
func (p *Place) SetLongitude(v float32) *Place {
	data, _ := json.Marshal(v)
	(*ld.Node)(p).SetNodes(as.Longitude, ld.Node{Value: data, Type: []string{xmlschema.TypeFloat}})
	return p
}

// SetLongitudeRaw sets the value in [as.Longitude].
func (p *Place) SetLongitudeRaw(v jsontext.Value) *Place {
	(*ld.Node)(p).SetNodes(as.Longitude, ld.Node{Value: v, Type: []string{xmlschema.TypeFloat}})
	return p
}

// See [Object.GetName].
func (p *Place) GetName() iter.Seq[*Localised] {
	return (*Object)(p).GetName()
}

// See [Object.AddName].
func (p *Place) AddName(ls ...Localised) *Place {
	(*Object)(p).AddName(ls...)
	return p
}

// See [Object.GetURL].
func (p *Place) GetURL() string {
	return (*Object)(p).GetURL()
}

// See [Object.SetURL].
func (p *Place) SetURL(url string) *Place {
	(*Object)(p).SetURL(url)
	return p
}
