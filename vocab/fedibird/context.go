// Package fedibird contains terms for the Fedibird namespace.
package fedibird

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://fedibird.com/ns#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "fedibird"

const (
	QuoteUri = Namespace + "quoteUri"
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
