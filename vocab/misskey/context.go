// Package misskey contains terms for the Misskey namespace.
package misskey

import "strings"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://misskey-hub.net/ns#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "misskey"

const (
	Content                      = Namespace + "_misskey_content"
	FollowedMessage              = Namespace + "_misskey_followedMessage"
	IsCat                        = Namespace + "isCat"
	License                      = Namespace + "_misskey_license"
	MakeNotesFollowersOnlyBefore = Namespace + "_misskey_makeNotesFollowersOnlyBefore"
	MakeNotesHiddenBefore        = Namespace + "_misskey_makeNotesHiddenBefore"
	Quote                        = Namespace + "_misskey_quote"
	Reaction                     = Namespace + "_misskey_reaction"
	RequireSigninToViewContents  = Namespace + "_misskey_requireSigninToViewContents"
	Summary                      = Namespace + "_misskey_summary"
	Talk                         = Namespace + "_misskey_talk"
	Votes                        = Namespace + "_misskey_votes"
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
