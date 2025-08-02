// Package vcard contains terms for the vCard namespace.
package vcard

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://www.w3.org/2006/vcard/ns#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "vcard"

func CompactIRI(iri string) string {
	return Prefix + `:` + Term(iri)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) any {
	// Already included in ActivityStreams context which is always present.
	return nil
}
