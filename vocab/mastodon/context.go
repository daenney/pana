// Package mastodon contains terms for the toot namespace.
package mastodon

import (
	"strings"

	ld "sourcery.dny.nu/longdistance"
)

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://joinmastodon.org/ns#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "toot"

const (
	AttributionDomains = Namespace + "attributionDomains"
	Blurhash           = Namespace + "blurhash"
	Discoverable       = Namespace + "discoverable"
	Featured           = Namespace + "featured"
	FeaturedTags       = Namespace + "featuredTags"
	FocalPoint         = Namespace + "focalPoint"
	Indexable          = Namespace + "indexable"
	Memorial           = Namespace + "memorial"
	Suspended          = Namespace + "suspended"
	VotersCount        = Namespace + "votersCount"
	TypeEmoji          = Namespace + "Emoji"
	TypeIdentityProof  = Namespace + "IdentityProof"
)

func CompactIRI(iri string) string {
	return Prefix + ":" + Term(iri)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) any {
	term := Term(iri)
	short := CompactIRI(iri)

	res := map[string]any{
		Prefix: Namespace,
	}

	switch iri {
	case AttributionDomains, Featured, FeaturedTags:
		res[term] = map[string]any{
			ld.KeywordID:   short,
			ld.KeywordType: ld.KeywordID,
		}
	case FocalPoint:
		res[term] = map[string]string{
			ld.KeywordID:        short,
			ld.KeywordContainer: ld.KeywordList,
		}
	default:
		res[term] = short
	}

	return res
}
