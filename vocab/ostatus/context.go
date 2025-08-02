// Package ostatus contains terms for the Ostatus.org namespace.
package ostatus

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://ostatus.org#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "ostatus"

const (
	AtomURI          = Namespace + "atomUri"
	Conversation     = Namespace + "conversation"
	InReplyToAtomURI = Namespace + "inReplyToAtomUri"
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
