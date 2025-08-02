// Package permsv1 contains terms for the W3ID Permissions namespace.
package permsv1

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/permissions#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "perm"

const (
	// AccessControl is an IRI, either as a string or as an object with an
	// id property.
	AccessControl = Namespace + "accessControl"
	// WritePermission is an IRI, either as a string or as an object with an
	// id property.
	WritePermission = Namespace + "writePermission"
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
