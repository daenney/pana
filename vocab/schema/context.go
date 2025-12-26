// Package schema contains terms for the Schema.org namespace.
package schema

import (
	"strings"
)

// Namespace is the IRI prefix used for terms defined in this namespace.
//
// This uses the fragment-suffixed namespace, as even though that's incorrect it
// is what's in use across ActivityPub implementations.
const Namespace = "http://schema.org#"

// Prefix is the canonical shorthand for [Namespace].
const Prefix = "schema"

const (
	About             = Namespace + "about"
	Address           = Namespace + "address"
	AddressCountry    = Namespace + "addressCountry"
	AddressLocality   = Namespace + "addressLocality"
	AddressRegion     = Namespace + "addressRegion"
	Description       = Namespace + "description"
	Email             = Namespace + "email"
	FamilyName        = Namespace + "familyName"
	GivenNaame        = Namespace + "givenName"
	Image             = Namespace + "image"
	Member            = Namespace + "member"
	MemberOf          = Namespace + "memberOf"
	Name              = Namespace + "name"
	PostalCode        = Namespace + "postalCode"
	StreetAddress     = Namespace + "streetAddress"
	TypePropertyValue = Namespace + "PropertyValue"
	TypeOrganization  = Namespace + "Organization"
	TypePerson        = Namespace + "Person"
	TypePostalAddress = Namespace + "PostalAddress"
	URL               = Namespace + "url"
	Value             = Namespace + "value"
)

func CompactIRI(iri string) string {
	return Prefix + `:` + Term(iri)
}

func Term(iri string) string {
	return strings.TrimPrefix(iri, Namespace)
}

func TermDefForIRI(iri string) map[string]any {
	term := Term(iri)

	return map[string]any{
		Prefix: Namespace,
		term:   Prefix + ":" + term,
	}
}
