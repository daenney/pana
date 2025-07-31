// Package xmlschema contains terms for the XML Schema namespace.
package xmlschema

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://www.w3.org/2001/XMLSchema#"

const (
	// TypeBoolean is a possible value for the type property.
	TypeBoolean = Namespace + "boolean"
	// TypeDateTime is a possible value for the type property.
	TypeDateTime = Namespace + "dateTime"
	// TypeDuration is a possible value for the type property.
	TypeDuration = Namespace + "duration"
	// TypeFloat is a possible value for the type property.
	TypeFloat = Namespace + "float"
	// TypeNonNegativeInteger is a possible value for the type property.
	TypeNonNegativeInteger = Namespace + "nonNegativeInteger"
)
