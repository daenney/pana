package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/litepub"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Note is the ActivityStreams Note type.
type Note Object

// NewNote initialises a new Note.
func NewNote() *Note {
	return &Note{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeNote},
	}
}

// Build finalises the Note.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (n *Note) Build() Any {
	return Any(*n)
}

// See [Object.GetID].
func (n *Note) GetID() string {
	return (*Object)(n).GetID()
}

// See [Object.SetID].
func (n *Note) SetID(id string) *Note {
	(*Object)(n).SetID(id)
	return n
}

// See [Object.GetType].
func (n *Note) GetType() string {
	return (*Object)(n).GetType()
}

// SetType sets the type to [as.Note].
func (n *Note) SetType() {
	n.Type = []string{as.TypeNote}
}

// See [Activity.GetActor].
func (n *Note) GetActor() iter.Seq[string] {
	return (*Activity)(n).GetActor()
}

// See [Activity.AddActor].
func (n *Note) AddActor(ids ...string) *Note {
	(*Activity)(n).AddActor(ids...)
	return n
}

// See [Object.GetAtomURI].
func (n *Note) GetAtomURI() string {
	return (*Object)(n).GetAtomURI()
}

// See [Object.SetAtomURI].
func (n *Note) SetAtomURI(uri string) *Note {
	(*Object)(n).SetAtomURI(uri)
	return n
}

// See [Object.GetAttachment].
func (n *Note) GetAttachment() iter.Seq[*Any] {
	return (*Object)(n).GetAttachment()
}

// See [Object.AddAttachment].
func (n *Note) AddAttachment(atch ...Any) *Note {
	(*Object)(n).AddAttachment(atch...)
	return n
}

// See [Object.GetAttributedTo].
func (n *Note) GetAttributedTo() iter.Seq[string] {
	return (*Object)(n).GetAttributedTo()
}

// See [Object.AddAttributedTo].
func (n *Note) AddAttributedTo(ids ...string) *Note {
	(*Object)(n).AddAttributedTo(ids...)
	return n
}

// See [Object.GetCc].
func (n *Note) GetCc() iter.Seq[string] {
	return (*Object)(n).GetCc()
}

// See [Object.AddCc].
func (n *Note) AddCc(ids ...string) *Note {
	(*Object)(n).AddCc(ids...)
	return n
}

// See [Object.GetContent].
func (n *Note) GetContent() iter.Seq[*Localised] {
	return (*Object)(n).GetContent()
}

// See [Object.AddContent].
func (n *Note) AddContent(ls ...Localised) *Note {
	(*Object)(n).AddContent(ls...)
	return n
}

// See [Object.GetContext].
func (n *Note) GetContext() string {
	return (*Object)(n).GetContext()
}

// See [Object.SetContext].
func (n *Note) SetContext(id string) *Note {
	(*Object)(n).SetContext(id)
	return n
}

// See [Object.GetConversation].
func (n *Note) GetConversation() string {
	return (*Object)(n).GetConversation()
}

// See [Object.SetConversation].
func (n *Note) SetConversation(id string) *Note {
	(*Object)(n).SetConversation(id)
	return n
}

// GetDirectMessage returns the value in [litepub.DirectMessage].
//
// It returns false if the value is absent.
func (n *Note) GetDirectMessage() json.RawMessage {
	if nodes := (*ld.Node)(n).GetNodes(litepub.DirectMessage); len(nodes) == 1 {
		return nodes[0].Value
	}

	return json.RawMessage(`false`)
}

// SetDirectMessage sets the value in [litepub.DirectMessage].
func (n *Note) SetDirectMessage(v json.RawMessage) *Note {
	(*ld.Node)(n).SetNodes(litepub.DirectMessage, ld.Node{Value: v})
	return n
}

// See [Object.GetInReplyTo].
func (n *Note) GetInReplyTo() string {
	return (*Object)(n).GetInReplyTo()
}

// See [Object.SetInReplyTo].
func (n *Note) SetInReplyTo(id string) *Note {
	(*Object)(n).SetInReplyTo(id)
	return n
}

// See [Object.GetInReplyToAtomURI].
func (n *Note) GetInReplyToAtomURI() string {
	return (*Object)(n).GetInReplyToAtomURI()
}

// See [Object.SetInReplyToAtomURI].
func (n *Note) SetInReplyToAtomURI(id string) *Note {
	(*Object)(n).SetInReplyToAtomURI(id)
	return n
}

// See [Object.GetLikes].
func (n *Note) GetLikes() *Collection {
	return (*Object)(n).GetLikes()
}

// See [Object.SetLikes].
func (n *Note) SetLikes(c Collection) *Note {
	(*Object)(n).SetLikes(c)
	return n
}

// See [Object.GetPublished].
func (n *Note) GetPublished() json.RawMessage {
	return (*Object)(n).GetPublished()
}

// See [Object.SetPublished].
func (n *Note) SetPublished(v json.RawMessage) *Note {
	(*Object)(n).SetPublished(v)
	return n
}

// See [Object.GetReplies].
func (n *Note) GetReplies() *Collection {
	return (*Object)(n).GetReplies()
}

// See [Object.SetReplies].
func (n *Note) SetReplies(c Collection) *Note {
	(*Object)(n).SetReplies(c)
	return n
}

// See [Object.GetSensitive].
func (n *Note) GetSensitive() json.RawMessage {
	return (*Object)(n).GetSensitive()
}

// See [Object.SetSensitive].
func (n *Note) SetSensitive(v json.RawMessage) *Note {
	(*Object)(n).SetSensitive(v)
	return n
}

// See [Object.GetShares].
func (n *Note) GetShares() *Collection {
	return (*Object)(n).GetShares()
}

// See [Object.SetShares].
func (n *Note) SetShares(c Collection) *Note {
	(*Object)(n).SetShares(c)
	return n
}

// See [Object.GetSummary].
func (n *Note) GetSummary() iter.Seq[*Localised] {
	return (*Object)(n).GetSummary()
}

// See [Object.AddSummary].
func (n *Note) AddSummary(ls ...Localised) *Note {
	(*Object)(n).AddSummary(ls...)
	return n
}

// See [Object.GetTag].
func (n *Note) GetTag() iter.Seq[*Any] {
	return (*Object)(n).GetTag()
}

// See [Object.AddTag].
func (n *Note) AddTag(tags ...Any) *Note {
	(*Object)(n).AddTag(tags...)
	return n
}

// See [Object.GetTo].
func (n *Note) GetTo() iter.Seq[string] {
	return (*Object)(n).GetTo()
}

// See [Object.AddTo].
func (n *Note) AddTo(ids ...string) *Note {
	(*Object)(n).AddTo(ids...)
	return n
}

// See [Object.GetUpdated].
func (n *Note) GetUpdated() json.RawMessage {
	return (*Object)(n).GetUpdated()
}

// See [Object.SetUpdated].
func (n *Note) SetUpdated(v json.RawMessage) *Note {
	(*Object)(n).SetUpdated(v)
	return n
}

// See [Object.GetURL].
func (n *Note) GetURL() string {
	return (*Object)(n).GetURL()
}

// See [Object.SetURL].
func (n *Note) SetURL(id string) *Note {
	(*Object)(n).SetURL(id)
	return n
}
