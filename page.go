package pana

import (
	"encoding/json/jsontext"
	"iter"
	"time"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Page is the ActivityStreams Page type.
type Page Object

// NewPage initialises a new Page.
func NewPage() *Page {
	return &Page{
		Properties: make(ld.Properties),
		Type:       []string{as.TypePage},
	}
}

// Build finalises the Page.
func (p *Page) Build() Page {
	return *p
}

// See [Object.GetID].
func (p *Page) GetID() string {
	return (*Object)(p).GetID()
}

// See [Object.SetID].
func (p *Page) SetID(id string) *Page {
	(*Object)(p).SetID(id)
	return p
}

// See [Object.GetType].
func (p *Page) GetType() string {
	return (*Object)(p).GetType()
}

// See [Object.SetType].
func (p *Page) SetType(typ string) *Page {
	(*Object)(p).SetType(typ)
	return p
}

// See [Object.GetAttachment].
func (p *Page) GetAttachment() iter.Seq[*Any] {
	return (*Object)(p).GetAttachment()
}

// See [Object.AddAttachment].
func (p *Page) AddAttachment[T Document | Audio | PropertyValue | ld.Node](atch ...T) *Page {
	(*Object)(p).AddAttachment(atch...)
	return p
}

// See [Object.GetAttributedTo].
func (p *Page) GetAttributedTo() iter.Seq[string] {
	return (*Object)(p).GetAttributedTo()
}

// See [Object.AddAttributedTo].
func (p *Page) AddAttributedTo(ids ...string) *Page {
	(*Object)(p).AddAttributedTo(ids...)
	return p
}

// See [Object.GetAudience].
func (p *Page) GetAudience() iter.Seq[string] {
	return (*Object)(p).GetAudience()
}

// See [Object.AddAudience].
func (p *Page) AddAudience(ids ...string) *Page {
	(*Object)(p).AddAudience(ids...)
	return p
}

// See [Object.GetCc].
func (p *Page) GetCc() iter.Seq[string] {
	return (*Object)(p).GetCc()
}

// See [Object.AddCc].
func (p *Page) AddCc(ids ...string) *Page {
	(*Object)(p).AddCc(ids...)
	return p
}

// See [Object.GetContent].
func (p *Page) GetContent() iter.Seq[*Localised] {
	return (*Object)(p).GetContent()
}

// See [Object.AddContent].
func (p *Page) AddContent(ls ...Localised) *Page {
	(*Object)(p).AddContent(ls...)
	return p
}

// See [Object.GetContext].
func (p *Page) GetContext() string {
	return (*Object)(p).GetContext()
}

// See [Object.SetContext].
func (p *Page) SetContext(id string) *Page {
	(*Object)(p).SetContext(id)
	return p
}

// See [Object.GetIcon].
func (p *Page) GetIcon() *Icon {
	return (*Object)(p).GetIcon()
}

// See [Object.SetIcon].
func (p *Page) SetIcon(img Icon) *Page {
	(*Object)(p).SetIcon(img)
	return p
}

// See [Object.GetImage].
func (p *Page) GetImage() *Image {
	return (*Object)(p).GetImage()
}

// See [Object.SetImage].
func (p *Page) SetImage(img Image) *Page {
	(*Object)(p).SetImage(img)
	return p
}

// See [Object.GetLikes].
func (p *Page) GetLikes() *Collection {
	return (*Object)(p).GetLikes()
}

// See [Object.SetLikes].
func (p *Page) SetLikes(c Collection) *Page {
	(*Object)(p).SetLikes(c)
	return p
}

// See [Object.GetMediaType].
func (p *Page) GetMediaType() jsontext.Value {
	return (*Object)(p).GetMediaType()
}

// See [Object.SetMediaType].
func (p *Page) SetMediaType(v string) *Page {
	(*Object)(p).SetMediaType(v)
	return p
}

// See [Object.SetMediaTypeRaw].
func (p *Page) SetMediaTypeRaw(v jsontext.Value) *Page {
	(*Object)(p).SetMediaTypeRaw(v)
	return p
}

// See [Object.GetName].
func (p *Page) GetName() iter.Seq[*Localised] {
	return (*Object)(p).GetName()
}

// See [Object.AddName].
func (p *Page) AddName(ls ...Localised) *Page {
	(*Object)(p).AddName(ls...)
	return p
}

// See [Object.GetPublished].
func (p *Page) GetPublished() jsontext.Value {
	return (*Object)(p).GetPublished()
}

// See [Object.SetPublished].
func (p *Page) SetPublished(v time.Time) *Page {
	(*Object)(p).SetPublished(v)
	return p
}

// See [Object.SetPublishedRaw].
func (p *Page) SetPublishedRaw(v jsontext.Value) *Page {
	(*Object)(p).SetPublishedRaw(v)
	return p
}

// See [Object.GetReplies].
func (p *Page) GetReplies() *Collection {
	return (*Object)(p).GetReplies()
}

// See [Object.SetReplies].
func (p *Page) SetReplies(c Collection) *Page {
	(*Object)(p).SetReplies(c)
	return p
}

// See [Object.GetSensitive].
func (p *Page) GetSensitive() jsontext.Value {
	return (*Object)(p).GetSensitive()
}

// See [Object.SetSensitive].
func (p *Page) SetSensitive(v bool) *Page {
	(*Object)(p).SetSensitive(v)
	return p
}

// See [Object.SetSensitiveRaw].
func (p *Page) SetSensitiveRaw(v jsontext.Value) *Page {
	(*Object)(p).SetSensitiveRaw(v)
	return p
}

// See [Object.GetShares].
func (p *Page) GetShares() *Collection {
	return (*Object)(p).GetShares()
}

// See [Object.SetShares].
func (p *Page) SetShares(c Collection) *Page {
	(*Object)(p).SetShares(c)
	return p
}

// See [Note.GetSource].
func (p *Page) GetSource() *Source {
	return (*Note)(p).GetSource()
}

// See [Note.SetSource].
func (p *Page) SetSource(s Source) *Page {
	(*Note)(p).SetSource(s)
	return p
}

// See [Object.GetSummary].
func (p *Page) GetSummary() iter.Seq[*Localised] {
	return (*Object)(p).GetSummary()
}

// See [Object.AddSummary].
func (p *Page) AddSummary(ls ...Localised) *Page {
	(*Object)(p).AddSummary(ls...)
	return p
}

// See [Object.GetTag].
func (p *Page) GetTag() iter.Seq[*Any] {
	return (*Object)(p).GetTag()
}

// See [Object.AddTag].
func (p *Page) AddTag[T LinkTag | Emoji | ld.Node](tags ...T) *Page {
	(*Object)(p).AddTag(tags...)
	return p
}

// See [Object.GetTo].
func (p *Page) GetTo() iter.Seq[string] {
	return (*Object)(p).GetTo()
}

// See [Object.AddTo].
func (p *Page) AddTo(ids ...string) *Page {
	(*Object)(p).AddTo(ids...)
	return p
}

// See [Object.GetUpdated].
func (p *Page) GetUpdated() jsontext.Value {
	return (*Object)(p).GetUpdated()
}

// See [Object.SetUpdated].
func (p *Page) SetUpdated(v time.Time) *Page {
	(*Object)(p).SetUpdated(v)
	return p
}

// See [Object.SetUpdatedRaw].
func (p *Page) SetUpdatedRaw(v jsontext.Value) *Page {
	(*Object)(p).SetUpdatedRaw(v)
	return p
}

// See [Object.GetURL].
func (p *Page) GetURL() string {
	return (*Object)(p).GetURL()
}

// See [Object.SetURL].
func (p *Page) SetURL(url string) *Page {
	(*Object)(p).SetURL(url)
	return p
}
