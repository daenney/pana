package pana

import (
	"encoding/json/jsontext"
	"iter"
	"time"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Article is the ActivityStreams Article type.
type Article Object

// NewArticle initialises a new Article.
func NewArticle() *Article {
	return &Article{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeArticle},
	}
}

// Build finalises the Article.
func (a *Article) Build() Article {
	return *a
}

// See [Object.GetID].
func (a *Article) GetID() string {
	return (*Object)(a).GetID()
}

// See [Object.SetID].
func (a *Article) SetID(id string) *Article {
	(*Object)(a).SetID(id)
	return a
}

// See [Object.GetType].
func (a *Article) GetType() string {
	return (*Object)(a).GetType()
}

// See [Object.SetType].
func (a *Article) SetType(typ string) *Article {
	(*Object)(a).SetType(typ)
	return a
}

// See [Object.GetAttachment].
func (a *Article) GetAttachment() iter.Seq[*Any] {
	return (*Object)(a).GetAttachment()
}

// See [Object.AddAttachment].
func (a *Article) AddAttachment[T Document | Audio | PropertyValue | ld.Node](atch ...T) *Article {
	(*Object)(a).AddAttachment(atch...)
	return a
}

// See [Object.GetAttributedTo].
func (a *Article) GetAttributedTo() iter.Seq[string] {
	return (*Object)(a).GetAttributedTo()
}

// See [Object.AddAttributedTo].
func (a *Article) AddAttributedTo(ids ...string) *Article {
	(*Object)(a).AddAttributedTo(ids...)
	return a
}

// See [Object.GetAudience].
func (a *Article) GetAudience() iter.Seq[string] {
	return (*Object)(a).GetAudience()
}

// See [Object.AddAudience].
func (a *Article) AddAudience(ids ...string) *Article {
	(*Object)(a).AddAudience(ids...)
	return a
}

// See [Object.GetCc].
func (a *Article) GetCc() iter.Seq[string] {
	return (*Object)(a).GetCc()
}

// See [Object.AddCc].
func (a *Article) AddCc(ids ...string) *Article {
	(*Object)(a).AddCc(ids...)
	return a
}

// See [Object.GetContent].
func (a *Article) GetContent() iter.Seq[*Localised] {
	return (*Object)(a).GetContent()
}

// See [Object.AddContent].
func (a *Article) AddContent(ls ...Localised) *Article {
	(*Object)(a).AddContent(ls...)
	return a
}

// See [Object.GetContext].
func (a *Article) GetContext() string {
	return (*Object)(a).GetContext()
}

// See [Object.SetContext].
func (a *Article) SetContext(id string) *Article {
	(*Object)(a).SetContext(id)
	return a
}

// See [Object.GetConversation].
func (a *Article) GetConversation() string {
	return (*Object)(a).GetConversation()
}

// See [Object.SetConversation].
func (a *Article) SetConversation(id string) *Article {
	(*Object)(a).SetConversation(id)
	return a
}

// See [Object.GetGenerator].
func (a *Article) GetGenerator() *Any {
	return (*Object)(a).GetGenerator()
}

// See [Object.SetGenerator].
func (a *Article) SetGenerator[T node](gen T) *Article {
	(*Object)(a).SetGenerator(gen)
	return a
}

// See [Object.GetIcon].
func (a *Article) GetIcon() *Icon {
	return (*Object)(a).GetIcon()
}

// See [Object.SetIcon].
func (a *Article) SetIcon(img Icon) *Article {
	(*Object)(a).SetIcon(img)
	return a
}

// See [Object.GetImage].
func (a *Article) GetImage() *Image {
	return (*Object)(a).GetImage()
}

// See [Object.SetImage].
func (a *Article) SetImage(img Image) *Article {
	(*Object)(a).SetImage(img)
	return a
}

// See [Object.GetInReplyTo].
func (a *Article) GetInReplyTo() string {
	return (*Object)(a).GetInReplyTo()
}

// See [Object.SetInReplyTo].
func (a *Article) SetInReplyTo(id string) *Article {
	(*Object)(a).SetInReplyTo(id)
	return a
}

// See [Object.GetLikes].
func (a *Article) GetLikes() *Collection {
	return (*Object)(a).GetLikes()
}

// See [Object.SetLikes].
func (a *Article) SetLikes(c Collection) *Article {
	(*Object)(a).SetLikes(c)
	return a
}

// See [Object.GetLocation].
func (a *Article) GetLocation() *Place {
	return (*Object)(a).GetLocation()
}

// See [Object.SetLocation].
func (a *Article) SetLocation(p Place) *Article {
	(*Object)(a).SetLocation(p)
	return a
}

// See [Object.GetMediaType].
func (a *Article) GetMediaType() jsontext.Value {
	return (*Object)(a).GetMediaType()
}

// See [Object.SetMediaType].
func (a *Article) SetMediaType(v string) *Article {
	(*Object)(a).SetMediaType(v)
	return a
}

// See [Object.SetMediaTypeRaw].
func (a *Article) SetMediaTypeRaw(v jsontext.Value) *Article {
	(*Object)(a).SetMediaTypeRaw(v)
	return a
}

// See [Object.GetName].
func (a *Article) GetName() iter.Seq[*Localised] {
	return (*Object)(a).GetName()
}

// See [Object.AddName].
func (a *Article) AddName(ls ...Localised) *Article {
	(*Object)(a).AddName(ls...)
	return a
}

// See [Object.GetPreview].
func (a *Article) GetPreview() *Any {
	return (*Object)(a).GetPreview()
}

// See [Object.SetPreview].
func (a *Article) SetPreview[T node](preview T) *Article {
	(*Object)(a).SetPreview(preview)
	return a
}

// See [Object.GetPublished].
func (a *Article) GetPublished() jsontext.Value {
	return (*Object)(a).GetPublished()
}

// See [Object.SetPublished].
func (a *Article) SetPublished(v time.Time) *Article {
	(*Object)(a).SetPublished(v)
	return a
}

// See [Object.SetPublishedRaw].
func (a *Article) SetPublishedRaw(v jsontext.Value) *Article {
	(*Object)(a).SetPublishedRaw(v)
	return a
}

// See [Object.GetReplies].
func (a *Article) GetReplies() *Collection {
	return (*Object)(a).GetReplies()
}

// See [Object.SetReplies].
func (a *Article) SetReplies(c Collection) *Article {
	(*Object)(a).SetReplies(c)
	return a
}

// See [Object.GetSensitive].
func (a *Article) GetSensitive() jsontext.Value {
	return (*Object)(a).GetSensitive()
}

// See [Object.SetSensitive].
func (a *Article) SetSensitive(v bool) *Article {
	(*Object)(a).SetSensitive(v)
	return a
}

// See [Object.SetSensitiveRaw].
func (a *Article) SetSensitiveRaw(v jsontext.Value) *Article {
	(*Object)(a).SetSensitiveRaw(v)
	return a
}

// See [Object.GetShares].
func (a *Article) GetShares() *Collection {
	return (*Object)(a).GetShares()
}

// See [Object.SetShares].
func (a *Article) SetShares(c Collection) *Article {
	(*Object)(a).SetShares(c)
	return a
}

// See [Note.GetSource].
func (a *Article) GetSource() *Source {
	return (*Note)(a).GetSource()
}

// See [Note.SetSource].
func (a *Article) SetSource(s Source) *Article {
	(*Note)(a).SetSource(s)
	return a
}

// See [Object.GetSummary].
func (a *Article) GetSummary() iter.Seq[*Localised] {
	return (*Object)(a).GetSummary()
}

// See [Object.AddSummary].
func (a *Article) AddSummary(ls ...Localised) *Article {
	(*Object)(a).AddSummary(ls...)
	return a
}

// See [Object.GetTag].
func (a *Article) GetTag() iter.Seq[*Any] {
	return (*Object)(a).GetTag()
}

// See [Object.AddTag].
func (a *Article) AddTag[T LinkTag | Emoji | ld.Node](tags ...T) *Article {
	(*Object)(a).AddTag(tags...)
	return a
}

// See [Object.GetTo].
func (a *Article) GetTo() iter.Seq[string] {
	return (*Object)(a).GetTo()
}

// See [Object.AddTo].
func (a *Article) AddTo(ids ...string) *Article {
	(*Object)(a).AddTo(ids...)
	return a
}

// See [Object.GetUpdated].
func (a *Article) GetUpdated() jsontext.Value {
	return (*Object)(a).GetUpdated()
}

// See [Object.SetUpdated].
func (a *Article) SetUpdated(v time.Time) *Article {
	(*Object)(a).SetUpdated(v)
	return a
}

// See [Object.SetUpdatedRaw].
func (a *Article) SetUpdatedRaw(v jsontext.Value) *Article {
	(*Object)(a).SetUpdatedRaw(v)
	return a
}

// See [Object.GetURL].
func (a *Article) GetURL() string {
	return (*Object)(a).GetURL()
}

// See [Object.SetURL].
func (a *Article) SetURL(url string) *Article {
	(*Object)(a).SetURL(url)
	return a
}
