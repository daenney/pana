// Package ostatus contains terms for the Ostatus.org namespace.
package ostatus

import (
	"strings"

	ld "sourcery.dny.nu/longdistance"
)

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
	res := map[string]any{
		Prefix: Namespace,
	}

	switch iri {
	case Conversation:
		res[Term(iri)] = map[string]any{
			ld.KeywordID:   CompactIRI(iri),
			ld.KeywordType: ld.KeywordID,
		}
	default:
		res[Term(iri)] = CompactIRI(iri)
	}

	return res
}
