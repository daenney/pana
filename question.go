package pana

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"iter"
	"time"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/mastodon"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
)

// Question is the ActivityStreams Question object.
type Question IntransitiveActivity

// NewQuestion initialises a new Question.
func NewQuestion() *Question {
	return &Question{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeQuestion},
	}
}

// Build finalises the Question.
func (q *Question) Build() Question {
	return *q
}

// IsMultipleChoice checks if this is a multiple-choice question.
func (q *Question) IsMultipleChoice() bool {
	if q == nil {
		return false
	}

	return Has(q, as.AnyOf)
}

// See [Object.GetID].
func (q *Question) GetID() string {
	return (*Object)(q).GetID()
}

// See [Object.SetID].
func (q *Question) SetID(id string) *Question {
	(*Object)(q).SetID(id)
	return q
}

// See [Object.GetType].
func (q *Question) GetType() string {
	return (*Object)(q).GetType()
}

// SetType sets the type to [as.TypeQuestion].
func (q *Question) SetType() {
	q.Type = []string{as.TypeQuestion}
}

// GetVotersCount returns the value in [mastodon.VotersCount].
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (q *Question) GetVotersCount() jsontext.Value {
	if nodes := (*ld.Node)(q).GetNodes(mastodon.VotersCount); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetVotersCount sets the number in [mastodon.VotersCount].
func (q *Question) SetVotersCount(v uint64) *Question {
	data, _ := json.Marshal(v)
	(*ld.Node)(q).SetNodes(mastodon.VotersCount, ld.Node{Value: data})
	return q
}

// SetVotersCountRaw sets the value in [mastodon.VotersCount].
func (q *Question) SetVotersCountRaw(v jsontext.Value) *Question {
	(*ld.Node)(q).SetNodes(mastodon.VotersCount, ld.Node{Value: v})
	return q
}

// See [Object.GetEndTime].
func (q *Question) GetEndTime() jsontext.Value {
	return (*Object)(q).GetEndTime()
}

// See [Object.SetEndTime].
func (q *Question) SetEndTime(v time.Time) *Question {
	(*Object)(q).SetEndTime(v)
	return q
}

// See [Object.SetEndTimeRaw].
func (q *Question) SetEndTimeRaw(v jsontext.Value) *Question {
	(*Object)(q).SetEndTimeRaw(v)
	return q
}

// See [Object.GetSensitive].
func (q *Question) GetSensitive() jsontext.Value {
	return (*Object)(q).GetSensitive()
}

// See [Object.SetSensitive].
func (q *Question) SetSensitive(v bool) *Question {
	(*Object)(q).SetSensitive(v)
	return q
}

// See [Object.SetSensitiveRaw].
func (q *Question) SetSensitiveRaw(v jsontext.Value) *Question {
	(*Object)(q).SetSensitiveRaw(v)
	return q
}

// See [Object.GetUpdated].
func (q *Question) GetUpdated() jsontext.Value {
	return (*Object)(q).GetUpdated()
}

// See [Object.SetUpdated].
func (q *Question) SetUpdated(v time.Time) *Question {
	(*Object)(q).SetUpdated(v)
	return q
}

// See [Object.SetUpdatedRaw].
func (q *Question) SetUpdatedRaw(v jsontext.Value) *Question {
	(*Object)(q).SetUpdatedRaw(v)
	return q
}

// GetAnyOf gets the [Choice] in [as.AnyOf].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-oneof.
func (q *Question) GetAnyOf() iter.Seq[Choice] {
	return func(yield func(Choice) bool) {
		for _, n := range (*ld.Node)(q).GetNodes(as.AnyOf) {
			if !yield(Choice(n)) {
				return
			}
		}
	}
}

// AddAnyOf appends [Choice] to [as.AnyOf].
func (q *Question) AddAnyOf(chs ...Choice) *Question {
	addNodes((*ld.Node)(q), as.AnyOf, chs)
	return q
}

// GetOneOf gets the [Choice] in [as.OneOf].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-oneof.
func (q *Question) GetOneOf() iter.Seq[Choice] {
	return func(yield func(Choice) bool) {
		for _, n := range (*ld.Node)(q).GetNodes(as.OneOf) {
			if !yield(Choice(n)) {
				return
			}
		}
	}
}

// AddOneOf appends [Choice] to [as.OneOf].
func (q *Question) AddOneOf(chs ...Choice) *Question {
	addNodes((*ld.Node)(q), as.OneOf, chs)
	return q
}

// GetClosed returns the value in [as.Closed].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-closed.
func (q *Question) GetClosed() jsontext.Value {
	if nodes := (*ld.Node)(q).GetNodes(as.Closed); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetClosed sets the [time.Time] in [as.Closed].
func (q *Question) SetClosed(v time.Time) *Question {
	data, _ := json.Marshal(v.Format(time.RFC3339))
	(*ld.Node)(q).SetNodes(as.Closed, ld.Node{Value: data, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// SetClosedRawsets the value in [as.Closed].
func (q *Question) SetClosedRaw(v jsontext.Value) *Question {
	(*ld.Node)(q).SetNodes(as.Closed, ld.Node{Value: v, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// See [Object.GetName].
func (q *Question) GetName() iter.Seq[*Localised] {
	return (*Object)(q).GetName()
}

// See [Object.AddName].
func (q *Question) AddName(ls ...Localised) *Question {
	(*Object)(q).AddName(ls...)
	return q
}

// See [Object.GetAtomURI].
func (q *Question) GetAtomURI() string {
	return (*Object)(q).GetAtomURI()
}

// See [Object.SetAtomURI].
func (q *Question) SetAtomURI(uri string) *Question {
	(*Object)(q).SetAtomURI(uri)
	return q
}

// See [Object.GetAttachment].
func (q *Question) GetAttachment() iter.Seq[*Any] {
	return (*Object)(q).GetAttachment()
}

// See [Object.AddAttachment].
func (q *Question) AddAttachment[T Document | Audio | PropertyValue | ld.Node](atch ...T) *Question {
	(*Object)(q).AddAttachment(atch...)
	return q
}

// See [Object.GetAttributedTo].
func (q *Question) GetAttributedTo() iter.Seq[string] {
	return (*Object)(q).GetAttributedTo()
}

// See [Object.AddAttributedTo].
func (q *Question) AddAttributedTo(ids ...string) *Question {
	(*Object)(q).AddAttributedTo(ids...)
	return q
}

// See [Object.GetCc].
func (q *Question) GetCc() iter.Seq[string] {
	return (*Object)(q).GetCc()
}

// See [Object.AddCc].
func (q *Question) AddCc(ids ...string) *Question {
	(*Object)(q).AddCc(ids...)
	return q
}

// See [Object.GetContent].
func (q *Question) GetContent() iter.Seq[*Localised] {
	return (*Object)(q).GetContent()
}

// See [Object.AddContent].
func (q *Question) AddContent(ls ...Localised) *Question {
	(*Object)(q).AddContent(ls...)
	return q
}

// See [Object.GetConversation].
func (q *Question) GetConversation() string {
	return (*Object)(q).GetConversation()
}

// See [Object.SetConversation].
func (q *Question) SetConversation(id string) *Question {
	(*Object)(q).SetConversation(id)
	return q
}

// See [Object.GetInReplyTo].
func (q *Question) GetInReplyTo() string {
	return (*Object)(q).GetInReplyTo()
}

// See [Object.SetInReplyTo].
func (q *Question) SetInReplyTo(id string) *Question {
	(*Object)(q).SetInReplyTo(id)
	return q
}

// See [Object.GetInReplyToAtomURI].
func (q *Question) GetInReplyToAtomURI() string {
	return (*Object)(q).GetInReplyToAtomURI()
}

// See [Object.SetInReplyToAtomURI].
func (q *Question) SetInReplyToAtomURI(id string) *Question {
	(*Object)(q).SetInReplyToAtomURI(id)
	return q
}

// See [Object.GetLikes].
func (q *Question) GetLikes() *Collection {
	return (*Object)(q).GetLikes()
}

// See [Object.SetLikes].
func (q *Question) SetLikes(c Collection) *Question {
	(*Object)(q).SetLikes(c)
	return q
}

// See [Object.GetPublished].
func (q *Question) GetPublished() jsontext.Value {
	return (*Object)(q).GetPublished()
}

// See [Object.SetPublished].
func (q *Question) SetPublished(v time.Time) *Question {
	(*Object)(q).SetPublished(v)
	return q
}

// See [Object.SetPublishedRaw].
func (q *Question) SetPublishedRaw(v jsontext.Value) *Question {
	(*Object)(q).SetPublishedRaw(v)
	return q
}

// See [Object.GetReplies].
func (q *Question) GetReplies() *Collection {
	return (*Object)(q).GetReplies()
}

// See [Object.SetReplies].
func (q *Question) SetReplies(c Collection) *Question {
	(*Object)(q).SetReplies(c)
	return q
}

// See [Object.GetShares].
func (q *Question) GetShares() *Collection {
	return (*Object)(q).GetShares()
}

// See [Object.SetShares].
func (q *Question) SetShares(c Collection) *Question {
	(*Object)(q).SetShares(c)
	return q
}

// See [Object.GetSummary].
func (q *Question) GetSummary() iter.Seq[*Localised] {
	return (*Object)(q).GetSummary()
}

// See [Object.AddSummary].
func (q *Question) AddSummary(ls ...Localised) *Question {
	(*Object)(q).AddSummary(ls...)
	return q
}

// See [Object.GetTag].
func (q *Question) GetTag() iter.Seq[*Any] {
	return (*Object)(q).GetTag()
}

// See [Object.AddTag].
func (q *Question) AddTag[T LinkTag | Emoji | ld.Node](tags ...T) *Question {
	(*Object)(q).AddTag(tags...)
	return q
}

// See [Object.GetTo].
func (q *Question) GetTo() iter.Seq[string] {
	return (*Object)(q).GetTo()
}

// See [Object.AddTo].
func (q *Question) AddTo(ids ...string) *Question {
	(*Object)(q).AddTo(ids...)
	return q
}

// See [Object.GetURL].
func (q *Question) GetURL() string {
	return (*Object)(q).GetURL()
}

// See [Object.SetURL].
func (q *Question) SetURL(url string) *Question {
	(*Object)(q).SetURL(url)
	return q
}

// Choice represents a choice in a poll. It is a much more limited [Note].
type Choice Note

// NewChoice initialises a new Choice.
func NewChoice() *Choice {
	return &Choice{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeNote},
	}
}

// Build finalises the Choice.
func (c *Choice) Build() Choice {
	return *c
}

// See [Object.GetName].
func (c *Choice) GetName() iter.Seq[*Localised] {
	return (*Object)(c).GetName()
}

// See [Object.AddName].
func (c *Choice) AddName(ls ...Localised) *Choice {
	(*Object)(c).AddName(ls...)
	return c
}

// See [Object.GetReplies].
func (c *Choice) GetReplies() *Collection {
	return (*Object)(c).GetReplies()
}

// See [Object.SetReplies].
func (c *Choice) SetReplies(cl Collection) *Choice {
	(*Object)(c).SetReplies(cl)
	return c
}
