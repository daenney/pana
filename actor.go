package pana

import (
	"encoding/json"
	"iter"

	ld "code.dny.dev/longdistance"
	"code.dny.dev/pana/vocab/mastodon"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
	ldp "code.dny.dev/pana/vocab/w3/ldp"
	secv1 "code.dny.dev/pana/vocab/w3id/securityv1"
)

// Actor is the ActivityStreams Actor type.
//
// See https://www.w3.org/TR/activitypub/#actor-objects.
type Actor Object

type Application = Actor
type Group = Actor
type Organisation = Actor
type Organization = Organisation
type Person = Actor
type Service = Actor

// NewActor initialises a new Actor.
func NewActor() *Actor {
	return &Actor{
		Properties: make(ld.Properties),
		Type:       []string{as.TypePerson},
	}
}

// Build finalises the Actor.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (a *Actor) Build() Any {
	return Any(*a)
}

// See [Object.GetType].
func (a *Actor) GetType() string {
	return (*Object)(a).GetType()
}

// See [Object.SetType].
func (a *Actor) SetType(typ string) *Actor {
	(*Object)(a).SetType(typ)
	return a
}

// See [Object.GetName].
func (a *Actor) GetName() iter.Seq[*Localised] {
	return (*Object)(a).GetName()
}

// See [Object.AddName].
func (a *Actor) AddName(ls ...Localised) *Actor {
	(*Object)(a).AddName(ls...)
	return a
}

// GetMemorial returns the value in [mastodon.Memorial].
//
// It returns false if the property was absent.
func (a *Actor) GetMemorial() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(mastodon.Memorial); len(nodes) == 1 {
		return nodes[0].Value
	}

	return json.RawMessage(`false`)
}

// SetMemorial sets the value in [mastodon.Memorial].
func (a *Actor) SetMemorial(v json.RawMessage) *Actor {
	(*ld.Node)(a).SetNodes(mastodon.Memorial, ld.Node{Value: v})
	return a
}

// See [Object.GetSummary].
func (a *Actor) GetSummary() iter.Seq[*Localised] {
	return (*Object)(a).GetSummary()
}

// See [Object.AddSummary].
func (a *Actor) AddSummary(ls ...Localised) *Actor {
	(*Object)(a).AddSummary(ls...)
	return a
}

// GetIcon returns the [Image] in [as.Icon].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-icon.
func (a *Actor) GetIcon() *Image {
	if nodes := (*ld.Node)(a).GetNodes(as.Icon); len(nodes) == 1 {
		return (*Image)(&nodes[0])
	}

	return nil
}

// SetIcon sets the [Image] in [as.Icon].
func (a *Actor) SetIcon(img Image) *Actor {
	(*ld.Node)(a).SetNodes(as.Icon, ld.Node(img))
	return a
}

// GetImage returns the [Image] in [as.Image].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-image.
func (a *Actor) GetImage() *Image {
	if nodes := (*ld.Node)(a).GetNodes(as.Image); len(nodes) == 1 {
		return (*Image)(&nodes[0])
	}

	return nil
}

// SetImage sets the [Image] in [as.Image].
func (a *Actor) SetImage(img Image) *Actor {
	(*ld.Node)(a).SetNodes(as.Image, ld.Node(img))
	return a
}

// GetDiscoverable returns the value in [mastodon.Discoverable].
//
// When discoverable is absent this returns false. We treat discoverable as
// opt-in, not opt-out.
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (a *Actor) GetDiscoverable() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(mastodon.Discoverable); len(nodes) == 1 {
		return nodes[0].Value
	}

	return json.RawMessage(`false`)
}

// SetDiscoverable sets the value in [mastodon.Discoverable].
func (a *Actor) SetDiscoverable(v json.RawMessage) *Actor {
	(*ld.Node)(a).SetNodes(mastodon.Discoverable, ld.Node{Value: v})
	return a
}

// GetFeatured returns the [Collection] stored in [mastodon.Featured].
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (a *Actor) GetFeatured() *Collection {
	if nodes := (*ld.Node)(a).GetNodes(mastodon.Featured); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetFeatured sets the [Collection] in [mastodon.Featured].
func (a *Actor) SetFeatured(c Collection) *Actor {
	(*ld.Node)(a).SetNodes(mastodon.Featured, ld.Node(c))
	return a
}

// GetFeaturedTags returns the [Collection] stored in [mastodon.FeaturedTags].
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (a *Actor) GetFeaturedTags() *Collection {
	if nodes := (*ld.Node)(a).GetNodes(mastodon.FeaturedTags); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetFeaturedTags sets the [Collection] in [mastodon.FeaturedTags].
func (a *Actor) SetFeaturedTags(c Collection) *Actor {
	(*ld.Node)(a).SetNodes(mastodon.FeaturedTags, ld.Node(c))
	return a
}

// GetFollowers returns the [Collection] stored in [as.Followers].
//
// See https://www.w3.org/TR/activitypub/#followers.
func (a *Actor) GetFollowers() *Collection {
	if nodes := (*ld.Node)(a).GetNodes(as.Followers); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetFollowers sets the [Collection] in [as.Followers].
func (a *Actor) SetFollowers(c Collection) *Actor {
	(*ld.Node)(a).SetNodes(as.Followers, ld.Node(c))
	return a
}

// GetFollowing returns the [Collection] stored in [as.Following].
//
// See https://www.w3.org/TR/activitypub/#following.
func (a *Actor) GetFollowing() *Collection {
	if nodes := (*ld.Node)(a).GetNodes(as.Following); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetFollowing sets the [Collection] in [as.Following].
func (a *Actor) SetFollowing(c Collection) *Actor {
	(*ld.Node)(a).SetNodes(as.Following, ld.Node(c))
	return a
}

// GetLiked returns the [Collection] stored in [as.Liked].
//
// See https://www.w3.org/TR/activitypub/#liked.
func (a *Actor) GetLiked() *Collection {
	if nodes := (*ld.Node)(a).GetNodes(as.Liked); len(nodes) == 1 {
		return (*Collection)(&nodes[0])
	}

	return nil
}

// SetLiked sets the [Collection] in [as.Liked].
func (a *Actor) SetLiked(c Collection) *Actor {
	(*ld.Node)(a).SetNodes(as.Liked, ld.Node(c))
	return a
}

// GetEndpoints returns the [Endpoints] from [as.Endpoints].
func (a *Actor) GetEndpoints() *Endpoints {
	if nodes := (*ld.Node)(a).GetNodes(as.Endpoints); len(nodes) == 1 {
		return (*Endpoints)(&nodes[0])
	}

	return nil
}

// SetEndpoints sets the [Endpoints] in [as.Endpoints].
func (a *Actor) SetEndpoints(ep Endpoints) *Actor {
	(*ld.Node)(a).SetNodes(as.Endpoints, ld.Node(ep))
	return a
}

// GetIndox returns the URL in [ldp.Inbox].
func (a *Actor) GetInbox() string {
	if nodes := (*ld.Node)(a).GetNodes(ldp.Inbox); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetInbox sets the URL in [ldp.Inbox].
func (a *Actor) SetInbox(url string) *Actor {
	(*ld.Node)(a).SetNodes(ldp.Inbox, ld.Node{ID: url})
	return a
}

// GetIndox returns the URL in [as.Outbox].
func (a *Actor) GetOutbox() string {
	if nodes := (*ld.Node)(a).GetNodes(as.Outbox); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetOutbox sets the URL in [as.Outbox].
func (a *Actor) SetOutbox(url string) *Actor {
	(*ld.Node)(a).SetNodes(as.Outbox, ld.Node{ID: url})
	return a
}

// GetIndexable returns the value in [mastodon.Indexable].
//
// When discoverable is absent this returns false. We treat indexable as
// opt-in, not opt-out.
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (a *Actor) GetIndexable() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(mastodon.Indexable); len(nodes) == 1 {
		return nodes[0].Value
	}

	return json.RawMessage(`false`)
}

// SetIndexable sets the value in [mastodon.Indexable].
func (a *Actor) SetIndexable(v json.RawMessage) *Actor {
	(*ld.Node)(a).SetNodes(mastodon.Indexable, ld.Node{Value: v})
	return a
}

// GetManuallyApprovesFollowers returns the value in
// [as.ManuallyApprovesFollowers].
//
// See https://swicg.github.io/miscellany/#manuallyApprovesFollowers.
func (a *Actor) GetManuallyApprovesFollowers() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(as.ManuallyApprovesFollowers); len(nodes) == 1 {
		return nodes[0].Value
	}

	return json.RawMessage(`false`)
}

// SetManuallyApprovesFollowers sets the value in
// [as.ManuallyApprovesFollowers].
func (a *Actor) SetManuallyApprovesFollowers(v json.RawMessage) *Actor {
	(*ld.Node)(a).SetNodes(as.ManuallyApprovesFollowers, ld.Node{Value: v})
	return a
}

// GetMovedTo returns the URL from [as.MovedTo].
//
// See https://swicg.github.io/miscellany/#movedTo.
func (a *Actor) GetMovedTo() string {
	if nodes := (*ld.Node)(a).GetNodes(as.MovedTo); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetMovedTo sets the URL in [as.MovedTo].
func (a *Actor) SetMovedTo(url string) *Actor {
	(*ld.Node)(a).SetNodes(as.MovedTo, ld.Node{ID: url})
	return a
}

// GetPreferredUsername returns the value in [as.PreferredUsername].
//
// See https://www.w3.org/TR/activitypub/#actor-objects.
func (a *Actor) GetPreferredUsername() json.RawMessage {
	if nodes := (*ld.Node)(a).GetNodes(as.PreferredUsername); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetPreferredUsername sets the value in [as.PreferredUsername].
func (a *Actor) SetPreferredUsername(v json.RawMessage) *Actor {
	(*ld.Node)(a).SetNodes(as.PreferredUsername, ld.Node{Value: v})
	return a
}

// See [Object.GetPublished].
func (a *Actor) GetPublished() json.RawMessage {
	return (*Object)(a).GetPublished()
}

// See [Object.SetPublished].
func (a *Actor) SetPublished(v json.RawMessage) *Actor {
	(*Object)(a).SetPublished(v)
	return a
}

// GetPublickKey returns the [PublicKey] stored in [secv1.PublicKey].
func (a *Actor) GetPublicKey() *PublicKey {
	if nodes := (*ld.Node)(a).GetNodes(secv1.PublicKey); len(nodes) == 1 {
		return (*PublicKey)(&nodes[0])
	}

	return nil
}

// SetPublicKey sets the [PublicKey] in [secv1.PublicKey].
func (a *Actor) SetPublicKey(pk PublicKey) *Actor {
	(*ld.Node)(a).SetNodes(secv1.PublicKey, ld.Node(pk))
	return a
}

// See [Object.GetURL].
func (a *Actor) GetURL() string {
	return (*Object)(a).GetURL()
}

// See [Object.SetURL].
func (a *Actor) SetURL(url string) *Actor {
	(*Object)(a).SetURL(url)
	return a
}
