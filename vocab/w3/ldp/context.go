// Package ldp contains terms for the Linked Data Platform namespace.
package ldp

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://www.w3.org/ns/ldp#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "ldp"

const (
	// Inbox is an IRI, either as a string or as an object with an id property.
	Inbox = Namespace + "inbox"
)

func CompactIRI(iri string) string {
	return Prefix + `:` + strings.TrimPrefix(iri, Namespace)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) any {
	// Already included in ActivityStreams context which is always present.
	return nil
}
