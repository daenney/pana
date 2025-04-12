// Package litepub contains terms for the Litepub namespace.
package litepub

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

//go:embed litepub-0.1.jsonld
var ContextDocument1dot0 []byte

// IRI is the remote context IRI.
const IRI = "https://litepub.social/litepub/context.jsonld"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "http://litepub.social/ns#"

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
