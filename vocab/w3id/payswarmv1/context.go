// Package payswarmv1 contains terms for the W3ID Payswarm namespace.
package payswarmv1

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/payswarm#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "ps"

const (
	// PaymentProcessor is a string.
	PaymentProcessor = Namespace + "processor"
	// Preferences is an IRI, either as a string or as an object with an id
	// property.
	Preferences = Namespace + "preferences"
)

func CompactIRI(iri string) string {
	return Prefix + `:` + Term(iri)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) map[string]any {
	return map[string]any{
		Prefix:    Namespace,
		Term(iri): CompactIRI(iri),
	}
}
