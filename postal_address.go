package pana

import (
	"encoding/json/jsontext"
	"encoding/json/v2"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/schema"
)

// PostalAddress is the Schema.org PostalAddress type.
type PostalAddress Object

// NewPostalAddress initialises a new PostalAddress.
func NewPostalAddress() *PostalAddress {
	return &PostalAddress{
		Properties: make(ld.Properties),
		Type:       []string{schema.TypePostalAddress},
	}
}

// Build finalises the PostalAddress.
func (pa *PostalAddress) Build() PostalAddress {
	return *pa
}

// See [Object.GetType].
func (pa *PostalAddress) GetType() string {
	return (*Object)(pa).GetType()
}

// GetAddressCountry returns the value in [schema.AddressCountry].
func (pa *PostalAddress) GetAddressCountry() jsontext.Value {
	if nodes := (*ld.Node)(pa).GetNodes(schema.AddressCountry); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetAddressCountry sets the string in [schema.AddressCountry].
func (pa *PostalAddress) SetAddressCountry(v string) *PostalAddress {
	data, _ := json.Marshal(v)
	(*ld.Node)(pa).SetNodes(schema.AddressCountry, ld.Node{Value: data})
	return pa
}

// SetAddressCountryRaw sets the value in [schema.AddressCountry].
func (pa *PostalAddress) SetAddressCountryRaw(v jsontext.Value) *PostalAddress {
	(*ld.Node)(pa).SetNodes(schema.AddressCountry, ld.Node{Value: v})
	return pa
}

// GetAddressLocality returns the value in [schema.AddressLocality].
func (pa *PostalAddress) GetAddressLocality() jsontext.Value {
	if nodes := (*ld.Node)(pa).GetNodes(schema.AddressLocality); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetAddressLocality sets the string in [schema.AddressLocality].
func (pa *PostalAddress) SetAddressLocality(v string) *PostalAddress {
	data, _ := json.Marshal(v)
	(*ld.Node)(pa).SetNodes(schema.AddressLocality, ld.Node{Value: data})
	return pa
}

// SetAddressLocalityRaw sets the value in [schema.AddressLocality].
func (pa *PostalAddress) SetAddressLocalityRaw(v jsontext.Value) *PostalAddress {
	(*ld.Node)(pa).SetNodes(schema.AddressLocality, ld.Node{Value: v})
	return pa
}

// GetAddressRegion returns the value in [schema.AddressRegion].
func (pa *PostalAddress) GetAddressRegion() jsontext.Value {
	if nodes := (*ld.Node)(pa).GetNodes(schema.AddressRegion); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetAddressRegion sets the string in [schema.AddressRegion].
func (pa *PostalAddress) SetAddressRegion(v string) *PostalAddress {
	data, _ := json.Marshal(v)
	(*ld.Node)(pa).SetNodes(schema.AddressRegion, ld.Node{Value: data})
	return pa
}

// SetAddressRegionRaw sets the value in [schema.AddressRegion].
func (pa *PostalAddress) SetAddressRegionRaw(v jsontext.Value) *PostalAddress {
	(*ld.Node)(pa).SetNodes(schema.AddressRegion, ld.Node{Value: v})
	return pa
}

// GetPostalCode returns the value in [schema.PostalCode].
func (pa *PostalAddress) GetPostalCode() jsontext.Value {
	if nodes := (*ld.Node)(pa).GetNodes(schema.PostalCode); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetPostalCode sets the string in [schema.PostalCode].
func (pa *PostalAddress) SetPostalCode(v string) *PostalAddress {
	data, _ := json.Marshal(v)
	(*ld.Node)(pa).SetNodes(schema.PostalCode, ld.Node{Value: data})
	return pa
}

// SetPostalCodeRaw sets the value in [schema.PostalCode].
func (pa *PostalAddress) SetPostalCodeRaw(v jsontext.Value) *PostalAddress {
	(*ld.Node)(pa).SetNodes(schema.PostalCode, ld.Node{Value: v})
	return pa
}

// GetStreetAddress returns the value in [schema.StreetAddress].
func (pa *PostalAddress) GetStreetAddress() jsontext.Value {
	if nodes := (*ld.Node)(pa).GetNodes(schema.StreetAddress); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetStreetAddress sets the string in [schema.StreetAddress].
func (pa *PostalAddress) SetStreetAddress(v string) *PostalAddress {
	data, _ := json.Marshal(v)
	(*ld.Node)(pa).SetNodes(schema.StreetAddress, ld.Node{Value: data})
	return pa
}

// SetStreetAddressRaw sets the value in [schema.StreetAddress].
func (pa *PostalAddress) SetStreetAddressRaw(v jsontext.Value) *PostalAddress {
	(*ld.Node)(pa).SetNodes(schema.StreetAddress, ld.Node{Value: v})
	return pa
}
