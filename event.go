package pana

import (
	"encoding/json/jsontext"
	"iter"
	"time"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Event is the Activitystreams Event type.
type Event Object

// NewEvent initialises a new Event.
func NewEvent() *Event {
	return &Event{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeEvent},
	}
}

// Build finalises the Event.
func (e *Event) Build() Event {
	return *e
}

// See [Object.GetID].
func (e *Event) GetID() string {
	return (*Object)(e).GetID()
}

// See [Object.SetID].
func (e *Event) SetID(id string) *Event {
	(*Object)(e).SetID(id)
	return e
}

// See [Object.GetType].
func (e *Event) GetType() string {
	return (*Object)(e).GetType()
}

// See [Object.SetType].
func (e *Event) SetType(typ string) *Event {
	(*Object)(e).SetType(typ)
	return e
}

// See [Activity.GetActor].
func (e *Event) GetActor() iter.Seq[string] {
	return (*Activity)(e).GetActor()
}

// See [Activity.AddActor].
func (e *Event) AddActor(ids ...string) *Event {
	(*Activity)(e).AddActor(ids...)
	return e
}

// See [Object.GetAttachment].
func (e *Event) GetAttachment() iter.Seq[*Any] {
	return (*Object)(e).GetAttachment()
}

// See [Object.AddAttachment].
func (e *Event) AddAttachment[T Document | Audio | PropertyValue | ld.Node](atch ...T) *Event {
	(*Object)(e).AddAttachment(atch...)
	return e
}

// See [Object.GetAttributedTo].
func (e *Event) GetAttributedTo() iter.Seq[string] {
	return (*Object)(e).GetAttributedTo()
}

// See [Object.AddAttributedTo].
func (e *Event) AddAttributedTo(ids ...string) *Event {
	(*Object)(e).AddAttributedTo(ids...)
	return e
}

// See [Object.GetAudience].
func (e *Event) GetAudience() iter.Seq[string] {
	return (*Object)(e).GetAudience()
}

// See [Object.AddAudience].
func (e *Event) AddAudience(ids ...string) *Event {
	(*Object)(e).AddAudience(ids...)
	return e
}

// See [Object.GetCc].
func (e *Event) GetCc() iter.Seq[string] {
	return (*Object)(e).GetCc()
}

// See [Object.AddCc].
func (e *Event) AddCc(ids ...string) *Event {
	(*Object)(e).AddCc(ids...)
	return e
}

// See [Object.GetContent].
func (e *Event) GetContent() iter.Seq[*Localised] {
	return (*Object)(e).GetContent()
}

// See [Object.AddContent].
func (e *Event) AddContent(ls ...Localised) *Event {
	(*Object)(e).AddContent(ls...)
	return e
}

// See [Object.GetConversation].
func (e *Event) GetConversation() string {
	return (*Object)(e).GetConversation()
}

// See [Object.SetConversation].
func (e *Event) SetConversation(id string) *Event {
	(*Object)(e).SetConversation(id)
	return e
}

// See [Object.GetEndTime].
func (e *Event) GetEndTime() jsontext.Value {
	return (*Object)(e).GetEndTime()
}

// See [Object.SetEndTime].
func (e *Event) SetEndTime(v time.Time) *Event {
	(*Object)(e).SetEndTime(v)
	return e
}

// See [Object.SetEndTimeRaw].
func (e *Event) SetEndTimeRaw(v jsontext.Value) *Event {
	(*Object)(e).SetEndTimeRaw(v)
	return e
}

// See [Object.GetInReplyTo].
func (e *Event) GetInReplyTo() string {
	return (*Object)(e).GetInReplyTo()
}

// See [Object.SetInReplyTo].
func (e *Event) SetInReplyTo(id string) *Event {
	(*Object)(e).SetInReplyTo(id)
	return e
}

// See [Object.GetLocation].
func (e *Event) GetLocation() *Place {
	return (*Object)(e).GetLocation()
}

// See [Object.SetLocation].
func (e *Event) SetLocation(p Place) *Event {
	(*Object)(e).SetLocation(p)
	return e
}

// See [Object.GetMediaType].
func (e *Event) GetMediaType() jsontext.Value {
	return (*Object)(e).GetMediaType()
}

// See [Object.SetMediaType].
func (e *Event) SetMediaType(v string) *Event {
	(*Object)(e).SetMediaType(v)
	return e
}

// See [Object.SetMediaTypeRaw].
func (e *Event) SetMediaTypeRaw(v jsontext.Value) *Event {
	(*Object)(e).SetMediaTypeRaw(v)
	return e
}

// See [Object.GetName].
func (e *Event) GetName() iter.Seq[*Localised] {
	return (*Object)(e).GetName()
}

// See [Object.AddName].
func (e *Event) AddName(ls ...Localised) *Event {
	(*Object)(e).AddName(ls...)
	return e
}

// See [Object.GetPublished].
func (e *Event) GetPublished() jsontext.Value {
	return (*Object)(e).GetPublished()
}

// See [Object.SetPublished].
func (e *Event) SetPublished(v time.Time) *Event {
	(*Object)(e).SetPublished(v)
	return e
}

// See [Object.SetPublishedRaw].
func (e *Event) SetPublishedRaw(v jsontext.Value) *Event {
	(*Object)(e).SetPublishedRaw(v)
	return e
}

// See [Object.GetSensitive].
func (e *Event) GetSensitive() jsontext.Value {
	return (*Object)(e).GetSensitive()
}

// See [Object.SetSensitive].
func (e *Event) SetSensitive(v bool) *Event {
	(*Object)(e).SetSensitive(v)
	return e
}

// See [Object.SetSensitiveRaw].
func (e *Event) SetSensitiveRaw(v jsontext.Value) *Event {
	(*Object)(e).SetSensitiveRaw(v)
	return e
}

// See [Note.GetSource].
func (e *Event) GetSource() *Source {
	return (*Note)(e).GetSource()
}

// See [Note.SetSource].
func (e *Event) SetSource(s Source) *Event {
	(*Note)(e).SetSource(s)
	return e
}

// See [Object.GetStartTime].
func (e *Event) GetStartTime() jsontext.Value {
	return (*Object)(e).GetStartTime()
}

// See [Object.SetStartTime].
func (e *Event) SetStartTime(v time.Time) *Event {
	(*Object)(e).SetStartTime(v)
	return e
}

// See [Object.SetStartTimeRaw].
func (e *Event) SetStartTimeRaw(v jsontext.Value) *Event {
	(*Object)(e).SetStartTimeRaw(v)
	return e
}

// See [Object.GetSummary].
func (e *Event) GetSummary() iter.Seq[*Localised] {
	return (*Object)(e).GetSummary()
}

// See [Object.AddSummary].
func (e *Event) AddSummary(ls ...Localised) *Event {
	(*Object)(e).AddSummary(ls...)
	return e
}

// See [Object.GetTag].
func (e *Event) GetTag() iter.Seq[*Any] {
	return (*Object)(e).GetTag()
}

// See [Object.AddTag].
func (e *Event) AddTag[T LinkTag | Emoji | ld.Node](tags ...T) *Event {
	(*Object)(e).AddTag(tags...)
	return e
}

// See [Object.GetTo].
func (e *Event) GetTo() iter.Seq[string] {
	return (*Object)(e).GetTo()
}

// See [Object.AddTo].
func (e *Event) AddTo(ids ...string) *Event {
	(*Object)(e).AddTo(ids...)
	return e
}

// See [Object.GetUpdated].
func (e *Event) GetUpdated() jsontext.Value {
	return (*Object)(e).GetUpdated()
}

// See [Object.SetUpdated].
func (e *Event) SetUpdated(v time.Time) *Event {
	(*Object)(e).SetUpdated(v)
	return e
}

// See [Object.SetUpdatedRaw].
func (e *Event) SetUpdatedRaw(v jsontext.Value) *Event {
	(*Object)(e).SetUpdatedRaw(v)
	return e
}

// See [Object.GetURL].
func (e *Event) GetURL() string {
	return (*Object)(e).GetURL()
}

// See [Object.SetURL].
func (e *Event) SetURL(url string) *Event {
	(*Object)(e).SetURL(url)
	return e
}
