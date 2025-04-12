// Package activitystreams contains terms for the ActivityStreams namespace.
//
// It also includes extensions.
package activitystreams

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://www.w3.org/ns/activitystreams"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = IRI + "#"

// From the normative context definition.
const (
	// Accuracy is an xml:float, an IEEE single-precision 32-bit floating point
	// value equivalent to a Go float32.
	Accuracy = Namespace + "accuracy"
	// Actor is an IRI, either as a string or as an object with an id property.
	Actor = Namespace + "actor"
	// AlsoKnownAs is an IRI, either as a string or as an object with an id
	// property.
	AlsoKnownAs = Namespace + "alsoKnownAs"
	// Altitude is an xml:float, an IEEE single-precision 32-bit floating point
	// value equivalent to a Go float32.
	Altitude = Namespace + "altitude"
	// AnyOf is an IRI, either as a string or as an object with an id property.
	AnyOf = Namespace + "anyOf"
	// Attachment is an IRI, either as a string or as an object with an id
	// property.
	Attachment = Namespace + "attachment"
	// AttributedTo is an IRI, either as a string or as an object with an id
	// property.
	AttributedTo = Namespace + "attributedTo"
	// Audience is an IRI, either as a string or as an object with an id
	// property.
	Audience = Namespace + "audience"
	// Bcc is an IRI, either as a string or as an object with an id property.
	Bcc = Namespace + "bcc"
	// Bto is an IRI, either as a string or as an object with an id property.
	Bto = Namespace + "bto"
	// Cc is an IRI, either as a string or as an object with an id property.
	Cc = Namespace + "cc"
	// Closed is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Closed = Namespace + "closed"
	// Content is a string.
	Content = Namespace + "content"
	// Context is an IRI, either as a string or as an object with an id
	// property.
	Context = Namespace + "context"
	// Current is an IRI, either as a string or as an object with an id
	// property.
	Current = Namespace + "current"
	// Deleted is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Deleted = Namespace + "deleted"
	// Describes is an IRI, either as a string or as an object with an id
	// property.
	Describes = Namespace + "describes"
	// Duration is an xml:duration and does not have a Go equivalent, but can be
	// handled as a string.
	Duration = Namespace + "duration"
	// EndTime is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	EndTime = Namespace + "endTime"
	// Endpoints is an IRI, either as a string or as an object with an id
	// property.
	Endpoints = Namespace + "endpoints"
	// First is an IRI, either as a string or as an object with an id property.
	First = Namespace + "first"
	// Followers is an IRI, either as a string or as an object with an id
	// property.
	Followers = Namespace + "followers"
	// Following is an IRI, either as a string or as an object with an id
	// property.
	Following = Namespace + "following"
	// FormerType is an IRI, either as a string or as an object with an id
	// property.
	FormerType = Namespace + "formerType"
	// Generator is an IRI, either as a string or as an object with an id
	// property.
	Generator = Namespace + "generator"
	// Height is an xml:nonNegativeInteger, an "infinite size" integer. The XML
	// specification requires you to at least accept numbers with up to 16
	// digits. A Go uint64 may be sufficient depending on your usage. Remember
	// that you can only safely express up to 53-bit precision integers this way
	// since JSON treats integers as floats. For bigger values you'll need a
	// string.
	Height = Namespace + "height"
	// Href is an IRI, either as a string or as an object with an id property.
	Href = Namespace + "href"
	// Hreflang is a string.
	Hreflang = Namespace + "hreflang"
	// Icon is an IRI, either as a string or as an object with an id property.
	Icon = Namespace + "icon"
	// Image is an IRI, either as a string or as an object with an id property.
	Image = Namespace + "image"
	// InReplyTo is an IRI, either as a string or as an object with an id
	// property.
	InReplyTo = Namespace + "inReplyTo"
	// Instrument is an IRI, either as a string or as an object with an id
	// property.
	Instrument = Namespace + "instrument"
	// Items is an IRI, either as a string or as an object with an id property.
	Items = Namespace + "items"
	// Last is an IRI, either as a string or as an object with an id property.
	Last = Namespace + "last"
	// Latitude is an xml:float, an IEEE single-precision 32-bit floating point
	// value equivalent to a Go float32.
	Latitude = Namespace + "latitude"
	// Liked is an IRI, either as a string or as an object with an id property.
	Liked = Namespace + "liked"
	// Likes is an IRI, either as a string or as an object with an id property.
	Likes = Namespace + "likes"
	// Location is an IRI, either as a string or as an object with an id
	// property.
	Location = Namespace + "location"
	// Longitude is an xml:float, an IEEE single-precision 32-bit floating point
	// value equivalent to a Go float32.
	Longitude = Namespace + "longitude"
	// MediaType is a string.
	MediaType = Namespace + "mediaType"
	// Name is a string.
	Name = Namespace + "name"
	// Next is an IRI, either as a string or as an object with an id property.
	Next = Namespace + "next"
	// OauthAuthzEndpoint is an IRI, either as a string or as an object
	// with an id property.
	OauthAuthzEndpoint = Namespace + "oauthAuthorizationEndpoint"
	// OauthTokenEndpoint is an IRI, either as a string or as an object with an
	// id property.
	OauthTokenEndpoint = Namespace + "oauthTokenEndpoint"
	// Object is an IRI, either as a string or as an object with an id property.
	Object = Namespace + "object"
	// OneOf is an IRI, either as a string or as an object with an id property.
	OneOf = Namespace + "oneOf"
	// OrderedItems is an IRI, either as a string or as an object with an id
	// property.
	OrderedItems = Namespace + "items"
	// Origin is an IRI, either as a string or as an object with an id property.
	Origin = Namespace + "origin"
	// Outbox is an IRI, either as a string or as an object with an id property.
	Outbox = Namespace + "outbox"
	// PartOf is an IRI, either as a string or as an object with an id property.
	PartOf = Namespace + "partOf"
	// PreferredUsername is a string.
	PreferredUsername = Namespace + "preferredUsername"
	// Prev is an IRI, either as a string or as an object with an id property.
	Prev = Namespace + "prev"
	// Preview is an IRI, either as a string or as an object with an id
	// property.
	Preview = Namespace + "preview"
	// ProvideClientKey is an IRI, either as a string or as an object with an id
	// property.
	ProvideClientKey = Namespace + "provideClientKey"
	// ProxyURL is an IRI, either as a string or as an object with an id
	// property.
	ProxyURL = Namespace + "proxyUrl"
	// PublicCollection represents the Public collection, aka everyone. Use it
	// as a value in the [Audience], [To] and [Cc] poperties.
	PublicCollection = Namespace + "Public"
	// Published is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Published = Namespace + "published"
	// Radius is an xml:float, an IEEE single-precision 32-bit floating point
	// value equivalent to a Go float32.
	Radius = Namespace + "radius"
	// Rel is a string.
	Rel = Namespace + "rel"
	// Relationship is an IRI, either as a string or as an object with an id
	// property.
	Relationship = Namespace + "relationship"
	// RelationshipIsContact is a possible value for a relationship property.
	RelationshipIsContact = Namespace + "IsContact"
	// RelationshipIsFollowedBy is a possible value for a relationship property.
	RelationshipIsFollowedBy = Namespace + "IsFollowedBy"
	// RelationshipIsFollowing is a possible value for a relationship property.
	RelationshipIsFollowing = Namespace + "IsFollowing"
	// RelationshipIsMember is a possible value for a relationship property.
	RelationshipIsMember = Namespace + "IsMember"
	// Replies is an IRI, either as a string or as an object with an id
	// property.
	Replies = Namespace + "replies"
	// Result is an IRI, either as a string or as an object with an id property.
	Result = Namespace + "result"
	// SharedInbox is an IRI, either as a string or as an object with an id
	// property.
	SharedInbox = Namespace + "sharedInbox"
	// Shares is an IRI, either as a string or as an object with an id property.
	Shares = Namespace + "shares"
	// SignClientKey is an IRI, either as a string or as an object with an id
	// property.
	SignClientKey = Namespace + "signClientKey"
	// Source is a string.
	Source = Namespace + "source"
	// StartIndex is an xml:nonNegativeInteger, an "infinite size" integer. The
	// XML specification requires you to at least accept numbers with up to 16
	// digits. A Go uint64 may be sufficient depending on your usage. Remember
	// that you can only safely express up to 53-bit precision integers this way
	// since JSON treats integers as floats. For bigger values you'll need a
	// string.
	StartIndex = Namespace + "startIndex"
	// StartTime is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	StartTime = Namespace + "startTime"
	// Streams is an IRI, either as a string or as an object with an id
	// property.
	Streams = Namespace + "streams"
	// Subject is an IRI, either as a string or as an object with an id
	// property.
	Subject = Namespace + "subject"
	// Summary is a string.
	Summary = Namespace + "summary"
	// Tag is an IRI, either as a string or as an object with an id property.
	Tag = Namespace + "tag"
	// Target is an IRI, either as a string or as an object with an id property.
	Target = Namespace + "target"
	// To is an IRI, either as a string or as an object with an id property.
	To = Namespace + "to"
	// TotalItems is an xml:nonNegativeInteger, an "infinite size" integer. The
	// XML specification requires you to at least accept numbers with up to 16
	// digits. A Go uint64 may be sufficient depending on your usage. Remember
	// that you can only safely express up to 53-bit precision integers this way
	// since JSON treats integers as floats. For bigger values you'll need a
	// string.
	TotalItems = Namespace + "totalItems"
	// TypeAccept is a possible value for the type property.
	TypeAccept = Namespace + "Accept"
	// TypeActivity is a possible value for the type property.
	TypeActivity = Namespace + "Activity"
	// TypeAdd is a possible value for the type property.
	TypeAdd = Namespace + "Add"
	// TypeAnnounce is a possible value for the type property.
	TypeAnnounce = Namespace + "Announce"
	// TypeApplication is a possible value for the type property.
	TypeApplication = Namespace + "Application"
	// TypeArrive is a possible value for the type property.
	TypeArrive = Namespace + "Arrive"
	// TypeArticle is a possible value for the type property.
	TypeArticle = Namespace + "Article"
	// TypeAudio is a possible value for the type property.
	TypeAudio = Namespace + "Audio"
	// TypeBlock is a possible value for the type property.
	TypeBlock = Namespace + "Block"
	// TypeCollection is a possible value for the type property.
	TypeCollection = Namespace + "Collection"
	// TypeCollectionPage is a possible value for the type property.
	TypeCollectionPage = Namespace + "CollectionPage"
	// TypeCreate is a possible value for the type property.
	TypeCreate = Namespace + "Create"
	// TypeDelete is a possible value for the type property.
	TypeDelete = Namespace + "Delete"
	// TypeDislike is a possible value for the type property.
	TypeDislike = Namespace + "Dislike"
	// TypeDocument is a possible value for the type property.
	TypeDocument = Namespace + "Document"
	// TypeEvent is a possible value for the type property.
	TypeEvent = Namespace + "Event"
	// TypeFlag is a possible value for the type property.
	TypeFlag = Namespace + "Flag"
	// TypeFollow is a possible value for the type property.
	TypeFollow = Namespace + "Follow"
	// TypeGroup is a possible value for the type property.
	TypeGroup = Namespace + "Group"
	// TypeIgnore is a possible value for the type property.
	TypeIgnore = Namespace + "Ignore"
	// TypeImage is a possible value for the type property.
	TypeImage = Namespace + "Image"
	// TypeIntransitiveActivity is a possible value for the type property.
	TypeIntransitiveActivity = Namespace + "IntransitiveActivity"
	// TypeInvite is a possible value for the type property.
	TypeInvite = Namespace + "Invite"
	// TypeJoin is a possible value for the type property.
	TypeJoin = Namespace + "Join"
	// TypeLeave is a possible value for the type property.
	TypeLeave = Namespace + "Leave"
	// TypeLike is a possible value for the type property.
	TypeLike = Namespace + "Like"
	// TypeLink is a possible value for the type property.
	TypeLink = Namespace + "Link"
	// TypeListen is a possible value for the type property.
	TypeListen = Namespace + "Listen"
	// TypeMention is a possible value for the type property.
	TypeMention = Namespace + "Mention"
	// TypeMove is a possible value for the type property.
	TypeMove = Namespace + "Move"
	// TypeNote is a possible value for the type property.
	TypeNote = Namespace + "Note"
	// TypeObject is a possible value for the type property.
	TypeObject = Namespace + "Object"
	// TypeOffer is a possible value for the type property.
	TypeOffer = Namespace + "Offer"
	// TypeOrderedCollection is a possible value for the type property.
	TypeOrderedCollection = Namespace + "OrderedCollection"
	// TypeOrderedCollectionPage is a possible value for the type property.
	TypeOrderedCollectionPage = Namespace + "OrderedCollectionPage"
	// TypeOrganization is a possible value for the type property.
	TypeOrganization = Namespace + "Organization"
	// TypePage is a possible value for the type property.
	TypePage = Namespace + "Page"
	// TypePerson is a possible value for the type property.
	TypePerson = Namespace + "Person"
	// TypePlace is a possible value for the type property.
	TypePlace = Namespace + "Place"
	// TypeProfile is a possible value for the type property.
	TypeProfile = Namespace + "Profile"
	// TypeQuestion is a possible value for the type property.
	TypeQuestion = Namespace + "Question"
	// TypeRead is a possible value for the type property.
	TypeRead = Namespace + "Read"
	// TypeReject is a possible value for the type property.
	TypeReject = Namespace + "Reject"
	// TypeRelationship is a possible value for the type property.
	TypeRelationship = Namespace + "Relationship"
	// TypeRemove is a possible value for the type property.
	TypeRemove = Namespace + "Remove"
	// TypeService is a possible value for the type property.
	TypeService = Namespace + "Service"
	// TypeTentativeAccept is a possible value for the type property.
	TypeTentativeAccept = Namespace + "TentativeAccept"
	// TypeTentativeReject is a possible value for the type property.
	TypeTentativeReject = Namespace + "TentativeReject"
	// TypeTombstone is a possible value for the type property.
	TypeTombstone = Namespace + "Tombstone"
	// TypeTravel is a possible value for the type property.
	TypeTravel = Namespace + "Travel"
	// TypeUndo is a possible value for the type property.
	TypeUndo = Namespace + "Undo"
	// TypeUpdate is a possible value for the type property.
	TypeUpdate = Namespace + "Update"
	// TypeVideo is a possible value for the type property.
	TypeVideo = Namespace + "Video"
	// TypeView is a possible value for the type property.
	TypeView = Namespace + "View"
	// URL is an IRI, either as a string or as an object with an id property.
	URL = Namespace + "url"
	// Units is a string.
	Units = Namespace + "units"
	// Updated is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Updated = Namespace + "updated"
	// UploadMedia is an IRI, either as a string or as an object with an id
	// property.
	UploadMedia = Namespace + "uploadMedia"
	// Width is an xml:nonNegativeInteger, an "infinite size" integer. The XML
	// specification requires you to at least accept numbers with up to 16
	// digits. A Go uint64 may be sufficient depending on your usage. Remember
	// that you can only safely express up to 53-bit precision integers this way
	// since JSON treats integers as floats. For bigger values you'll need a
	// string.
	Width = Namespace + "width"
)

// Extensions: https://www.w3.org/wiki/Activity_Streams_extensions.
const (
	// ManuallyApprovesFollowers is an xml:boolean.
	ManuallyApprovesFollowers = Namespace + "manuallyApprovesFollowers"
	// MovedTo is an IRI, either as a string or as an object with an id
	// property.
	MovedTo = Namespace + "movedTo"
	// Sensitive is an xml:boolean.
	Sensitive = Namespace + "sensitive"
	// TypeHashtag is a possible value for the type property.
	TypeHashtag = Namespace + "Hashtag"
)
