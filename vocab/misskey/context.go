// Package misskey contains terms for the Misskey namespace.
package misskey

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://misskey-hub.net/ns#"

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
