package pana

import (
	ld "code.dny.dev/longdistance"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

// Endpoints from ActivityPub.
//
// Endpoints does not include support for [as.ProvideClientKey] or
// [as.SignClientKey] because Linked Data Signatures are a menace and should not
// be used. Use HTTP Signatures instead.
//
// See https://www.w3.org/TR/activitypub/#actor-objects.
type Endpoints ld.Node

// NewEndpoints initialises a new Endpoints.
func NewEndpoints() *Endpoints {
	return &Endpoints{
		Properties: make(ld.Properties, 4),
	}
}

// Build finalises the Endpoints.
func (ep *Endpoints) Build() Endpoints {
	return *ep
}

// GetProxyURL returns the URL in [as.ProxyURL].
func (ep *Endpoints) GetProxyURL() string {
	if nodes := (*ld.Node)(ep).GetNodes(as.ProxyURL); len(nodes) == 1 {
		return nodes[0].ID
	}
	return ""
}

// SetProxyURL sets the URL in [as.ProxyURL].
func (ep *Endpoints) SetProxyURL(url string) *Endpoints {
	(*ld.Node)(ep).SetNodes(as.ProxyURL, ld.Node{ID: url})
	return ep
}

// GetOauthAuthzEndpoint returns the URL in [as.OauthAuthzEndpoint].
func (ep *Endpoints) GetOauthAuthzEndpoint() string {
	if nodes := (*ld.Node)(ep).GetNodes(as.OauthAuthzEndpoint); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetOauthAuthzEndpoint sets the URL in [as.OauthAuthzEndpoint].
func (ep *Endpoints) SetOauthAuthzEndpoint(url string) *Endpoints {
	(*ld.Node)(ep).SetNodes(as.OauthAuthzEndpoint, ld.Node{ID: url})
	return ep
}

// GetOauthTokenEndpoint returns the URL in [as.OauthTokenEndpoint].
func (ep *Endpoints) GetOauthTokenEndpoint() string {
	if nodes := (*ld.Node)(ep).GetNodes(as.OauthTokenEndpoint); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetOauthTokenEndpoint sets the URL in [as.OauthTokenEndpoint].
func (ep *Endpoints) SetOauthTokenEndpoint(url string) *Endpoints {
	(*ld.Node)(ep).SetNodes(as.OauthTokenEndpoint, ld.Node{ID: url})
	return ep
}

// GetSharedInbox returns the URL in [as.SharedInbox].
func (ep *Endpoints) GetSharedInbox() string {
	if nodes := (*ld.Node)(ep).GetNodes(as.SharedInbox); len(nodes) == 1 {
		return nodes[0].ID
	}
	return ""
}

// SetSharedInbox sets the URL in [as.SharedInbox].
func (ep *Endpoints) SetSharedInbox(url string) *Endpoints {
	(*ld.Node)(ep).SetNodes(as.SharedInbox, ld.Node{ID: url})
	return ep
}
