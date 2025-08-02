// Package xmlschema contains terms for the XML Schema namespace.
package xmlschema

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://www.w3.org/2001/XMLSchema#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "xsd"

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

func CompactIRI(iri string) string {
	return Prefix + ":" + Term(iri)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) any {
	// Already included in ActivityStreams context which is always present.
	return nil
}
