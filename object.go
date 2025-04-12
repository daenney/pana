package pana

import (
	"encoding/json"
	"iter"

	ld "code.dny.dev/longdistance"
	"code.dny.dev/pana/vocab/ostatus"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Object is the ActivityStreams Object type.
type Object ld.Node

// NewObject initialises a new Object.
func NewObject() *Object {
	return &Object{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeObject},
	}
}

// Build finalises the Object.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (o *Object) Build() Any {
	return Any(*o)
}

// GetID returns the object ID from [ld.Node.ID].
//
// This is equivalent to the 'id' property in compacted form JSON or the '@id'
// property in expanded form.
//
// If an object doesn't have an ID but does have other properties it is an
// anonymous object (a blank node). This represents something that exists but is
// not named or referencable.
func (o *Object) GetID() string {
	return o.ID
}

// SetID sets the object ID in [ld.Node.ID].
func (o *Object) SetID(id string) *Object {
	o.ID = id
	return o
}

// GetType returns the first type from [ld.Node.Type].
//
// This is equivalent to the 'type' property in compacted form JSON or the
// '@type' property in expanded form.
//
// It will return the empty string if the object has no type.
//
// An object can have multiple types, but this is so uncommon on the fediverse
// that this method only returns the first type. You can check the size of
// [ld.Node.Type] if you'd like.
func (o *Object) GetType() string {
	if o.Type == nil {
		return ""
	}
	return o.Type[0]
}

// SetType sets the type in [ld.Node.Type].
func (o *Object) SetType(t string) *Object {
	o.Type = []string{t}
	return o
}

// GetCc returns the audience IDs from [as.Cc].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-cc.
func (o *Object) GetCc() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Cc) {
			if !yield(n.ID) {
				return
			}
		}
	}
}

// AddCc appends audience IDs to [as.Cc].
func (o *Object) AddCc(ids ...string) *Object {
	(*ld.Node)(o).AddNodes(as.Cc, toReference(ids...)...)
	return o
}

// GetContent returns the localised values in [as.Content].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-content.
func (o *Object) GetContent() iter.Seq[*Localised] {
	return func(yield func(*Localised) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Content) {
			if !yield((*Localised)(&n)) {
				return
			}
		}
	}
}

// AddContent adds values to [as.Content].
func (o *Object) AddContent(ls ...Localised) *Object {
	(*ld.Node)(o).AddNodes(as.Content, toLDNodes(ls...)...)
	return o
}

// GetContext returns the ID in [as.Context].
//
// This is not the JSON-LD [ld.KeywordContext].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-context.
func (o *Object) GetContext() string {
	if nodes := (*ld.Node)(o).GetNodes(as.Context); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetContext sets the ID in [as.Context].
func (o *Object) SetContext(id string) *Object {
	(*ld.Node)(o).SetNodes(as.Context, ld.Node{ID: id})
	return o
}

// GetTo returns the audience IDs from [as.To].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-to.
func (o *Object) GetTo() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.To) {
			if !yield(n.ID) {
				return
			}
		}
	}
}

// AddTo appends audience IDs to [as.To].
func (o *Object) AddTo(ids ...string) *Object {
	(*ld.Node)(o).AddNodes(as.To, toReference(ids...)...)
	return o
}

// GetPublished retrieves the value from [as.Published].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-published.
func (o *Object) GetPublished() json.RawMessage {
	if nodes := (*ld.Node)(o).GetNodes(as.Published); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetPublished sets the value in [as.Published].
func (o *Object) SetPublished(value json.RawMessage) *Object {
	(*ld.Node)(o).SetNodes(as.Published, ld.Node{Value: value})
	return o
}

// GetAttributedTo returns the actor IDs from [as.AttributedTo].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-attributedto.
func (o *Object) GetAttributedTo() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.AttributedTo) {
			if !yield(n.ID) {
				return
			}
		}
	}
}

// AddAttributedTo appends actor IDs to [as.AttributedTo].
func (o *Object) AddAttributedTo(ids ...string) *Object {
	(*ld.Node)(o).AddNodes(as.To, toReference(ids...)...)
	return o
}

// GetReplies returns the [Collection] stored in [as.Replies].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-replies.
func (o *Object) GetReplies() *Collection {
	if nodes := (*ld.Node)(o).GetNodes(as.Replies); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetReplies sets the [Collection] in [as.Replies].
func (o *Object) SetReplies(c Collection) *Object {
	(*ld.Node)(o).SetNodes(as.Replies, ld.Node(c))
	return o
}

// GetURL returns the URL in [as.URL].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-url.
func (o *Object) GetURL() string {
	if nodes := (*ld.Node)(o).GetNodes(as.URL); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetURL sets the URL in [as.URL].
func (o *Object) SetURL(url string) *Object {
	(*ld.Node)(o).SetNodes(as.URL, toReference(url)...)
	return o
}

// GetConversation returns the URI in [ostatus.Conversation].
//
// This is a [tag URI].
//
// [tag URI]: https://datatracker.ietf.org/doc/html/rfc4151
func (o *Object) GetConversation() string {
	if nodes := (*ld.Node)(o).GetNodes(ostatus.Conversation); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetConversation sets the URI in [ostatus.Conversation].
func (o *Object) SetConversation(uri string) *Object {
	(*ld.Node)(o).SetNodes(ostatus.Conversation, ld.Node{ID: uri})
	return o
}

// GetAtomUri returns the URI in [ostatus.AtomURI].
//
// [tag URI]: https://datatracker.ietf.org/doc/html/rfc4151
func (o *Object) GetAtomURI() string {
	if nodes := (*ld.Node)(o).GetNodes(ostatus.AtomURI); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetAtomUri sets the URI in [ostatus.AtomURI].
func (o *Object) SetAtomURI(uri string) *Object {
	(*ld.Node)(o).SetNodes(ostatus.AtomURI, ld.Node{ID: uri})
	return o
}

// GetSummary returns the localised values in [as.Summary].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-summary.
func (o *Object) GetSummary() iter.Seq[*Localised] {
	return func(yield func(*Localised) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Summary) {
			if !yield((*Localised)(&n)) {
				return
			}
		}
	}
}

// AddSummary adds values to [as.Summary].
func (o *Object) AddSummary(ls ...Localised) *Object {
	(*ld.Node)(o).AddNodes(as.Summary, toLDNodes(ls...)...)
	return o
}

// GetTag returns the values in [as.Tag].
//
// This returns Any because it can be an Emoji, a Hashtag, combinations of the
// two or more.
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-tag.
func (o *Object) GetTag() iter.Seq[*Any] {
	return func(yield func(*Any) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Tag) {
			if !yield((*Any)(&n)) {
				return
			}
		}
	}
}

// AddTag appends a tag to [as.Tag].
func (o *Object) AddTag(tag ...Any) *Object {
	(*ld.Node)(o).AddNodes(as.Tag, toLDNodes(tag...)...)
	return o
}

// GetInReplyTo returns the URL in [as.InReplyTo].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-inreplyto.
func (o *Object) GetInReplyTo() string {
	if nodes := (*ld.Node)(o).GetNodes(as.InReplyTo); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetInReplyTo sets the URL in [as.InReplyTo].
func (o *Object) SetInReplyTo(url string) *Object {
	(*ld.Node)(o).SetNodes(as.InReplyTo, ld.Node{ID: url})
	return o
}

// GetAttachment returns the attachments in [as.Attachment].
//
// This returns [Any] because is can be a Document, PropertyValue, a combination
// of both and more.
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-attachment.
func (o *Object) GetAttachment() iter.Seq[*Any] {
	return func(yield func(*Any) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Attachment) {
			if !yield((*Any)(&n)) {
				return
			}
		}
	}
}

// AddAttachment appends attachments to [as.Attachment].
func (o *Object) AddAttachment(atch ...Any) *Object {
	(*ld.Node)(o).AddNodes(as.Attachment, toLDNodes(atch...)...)
	return o
}

// GetSensitive returns the value from [as.Sensitive].
//
// See https://swicg.github.io/miscellany/#sensitive.
func (o *Object) GetSensitive() json.RawMessage {
	if nodes := (*ld.Node)(o).GetNodes(as.Sensitive); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetSensitive sets the value in [as.Sensitive].
func (o *Object) SetSensitive(v json.RawMessage) *Object {
	(*ld.Node)(o).SetNodes(as.Sensitive, ld.Node{Value: v})
	return o
}

// GetInReplyToAtomURI returns the URI in [ostatus.InReplyToAtomUri].
func (o *Object) GetInReplyToAtomURI() string {
	if nodes := (*ld.Node)(o).GetNodes(ostatus.InReplyToAtomURI); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetInReplyToAtomURI sets the URI in [ostatus.InReplyToAtomUri].
func (o *Object) SetInReplyToAtomURI(uri string) *Object {
	(*ld.Node)(o).SetNodes(ostatus.InReplyToAtomURI, ld.Node{ID: uri})
	return o
}

// GetUpdated returns the value in [as.Updated].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-endtime.
func (o *Object) GetUpdated() json.RawMessage {
	if nodes := (*ld.Node)(o).GetNodes(as.Updated); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetUpdated sets the value in [as.Updated].
func (o *Object) SetUpdated(v json.RawMessage) *Object {
	(*ld.Node)(o).SetNodes(as.Updated, ld.Node{Value: v})
	return o
}

// GetName returns the localised values in [as.Name].
func (o *Object) GetName() iter.Seq[*Localised] {
	return func(yield func(*Localised) bool) {
		for _, n := range (*ld.Node)(o).GetNodes(as.Name) {
			if !yield((*Localised)(&n)) {
				return
			}
		}
	}
}

// AddName appends localised values to [as.Name].
func (o *Object) AddName(ls ...Localised) *Object {
	(*ld.Node)(o).AddNodes(as.Name, toLDNodes(ls...)...)
	return o
}

// GetLikes returns the [Collection] stored in [as.Likes].
//
// See https://www.w3.org/TR/activitypub/#likes.
func (o *Object) GetLikes() *Collection {
	if nodes := (*ld.Node)(o).GetNodes(as.Likes); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetLikes sets the [Collection] in [as.Likes].
func (o *Object) SetLikes(c Collection) *Object {
	(*ld.Node)(o).SetNodes(as.Likes, ld.Node(c))
	return o
}

// GetShares returns the [Collection] stored in [as.Shares].
//
// See https://www.w3.org/TR/activitypub/#shares.
func (o *Object) GetShares() *Collection {
	if nodes := (*ld.Node)(o).GetNodes(as.Shares); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetShares sets the [Collection] in [as.Shares].
func (o *Object) SetShares(c Collection) *Object {
	(*ld.Node)(o).SetNodes(as.Shares, ld.Node(c))
	return o
}
