// Package terms contains terms for the Dublin Core Terms namespace.
package terms

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://purl.org/dc/terms/"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "dcterms"

const (
	Created     = Namespace + "created"
	Creator     = Namespace + "creator"
	Description = Namespace + "description"
	License     = Namespace + "license"
	Source      = Namespace + "source"
	Title       = Namespace + "title"
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
