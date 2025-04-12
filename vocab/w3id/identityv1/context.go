// Package identityv1 contains terms for the W3ID Identity namespace.
package identityv1

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://w3id.org/identity/v1"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/identity#"

const (
	// IdentityService is an IRI, either as a string or as an object with an
	// id property.
	IdentityService = Namespace + "identityService"
	// IDP is an IRI, either as a string or as an object with an
	// id property.
	IDP = Namespace + "idp"
	// TypeIdentity is a possible value for the type property.
	TypeIdentity = Namespace + "Identity"
)
