package pana

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/schema"
)

// PropertyValue is the Schema.org PropertyValue type.
type PropertyValue Object

// NewPropertyValue initialises a new PropertyValue.
func NewPropertyValue() *PropertyValue {
	return &PropertyValue{
		Properties: make(ld.Properties),
		Type:       []string{schema.TypePropertyValue},
	}
}

// Build finalises the PropertyValue.
func (pv *PropertyValue) Build() PropertyValue {
	return *pv
}

// See [Object.GetType].
func (pv *PropertyValue) GetType() string {
	return (*Object)(pv).GetType()
}

// See [Object.GetName].
func (pv *PropertyValue) GetName() iter.Seq[*Localised] {
	return (*Object)(pv).GetName()
}

// See [Object.AddName].
func (pv *PropertyValue) AddName(ls ...Localised) *PropertyValue {
	(*Object)(pv).AddName(ls...)
	return pv
}

// GetValue returns the value in [schema.Value].
func (pv *PropertyValue) GetValue() jsontext.Value {
	if nodes := (*ld.Node)(pv).GetNodes(schema.Value); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetValue sets the string in [schema.Value].
func (pv *PropertyValue) SetValue(v string) *PropertyValue {
	data, _ := json.Marshal(v)
	(*ld.Node)(pv).SetNodes(schema.Value, ld.Node{Value: data})
	return pv
}

// SetValueRaw sets the value in [schema.Value].
func (pv *PropertyValue) SetValueRaw(v jsontext.Value) *PropertyValue {
	(*ld.Node)(pv).SetNodes(schema.Value, ld.Node{Value: v})
	return pv
}
