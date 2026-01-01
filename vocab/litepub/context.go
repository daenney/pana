// Package litepub contains terms for the Litepub namespace.
package litepub

import (
	_ "embed"
	"strings"
)

//go:embed context.jsonld
var ContextDocument []byte

//go:embed litepub-0.1.jsonld
var ContextDocument0dot1 []byte

// IRI is the remote context IRI.
const IRI = "https://litepub.social/litepub/context.jsonld"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://litepub.social/ns#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "litepub"

const (
	Capabilities              = Namespace + "capabilities"
	DirectMessage             = Namespace + "directMessage"
	FormerRepresentations     = Namespace + "formerRepresentations"
	Invisible                 = Namespace + "invisible"
	ListMessage               = Namespace + "listMessage"
	OauthRegistrationEndpoint = Namespace + "oauthRegistrationEndpoint"
	TypeChatMessage           = Namespace + "ChatMessage"
	TypeEmojiReact            = Namespace + "EmojiReact"
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
