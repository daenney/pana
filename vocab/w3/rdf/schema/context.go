// Package schema contains terms for the RDF Schema namespace.
package schema

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://www.w3.org/2000/01/rdf-schema#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "rdfs"

const (
	// Comment is a string.
	Comment = Namespace + "comment"
	// Label is a string.
	Label = Namespace + "label"
)

func CompactIRI(iri string) string {
	return Prefix + `:` + strings.TrimPrefix(iri, Namespace)
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
