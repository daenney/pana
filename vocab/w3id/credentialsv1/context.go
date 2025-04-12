// Package credv1 contains terms for the W3ID Credentials namespace.
package credv1

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://w3id.org/credentials/v1"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/credentials#"

const (
	// Claim is an IRI, either as a string or as an object with an
	// id property.
	Claim = Namespace + "claim"
	// Credential is an IRI, either as a string or as an object with an
	// id property.
	Credential = Namespace + "credential"
	// Issued is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Issued = Namespace + "issued"
	// Issuer is an IRI, either as a string or as an object with an
	// id property.
	Issuer = Namespace + "issuer"
	// Recipient is an IRI, either as a string or as an object with an
	// id property.
	Recipient = Namespace + "recipient"
	// ReferenceID is a string.
	ReferenceID    = Namespace + "referenceId"
	TypeCredential = Namespace + "Credential"
	// TypeCryptographicKeyCredential is a possible value for the type property.
	TypeCryptographicKeyCredential = Namespace + "CryptographicKeyCredential"
)
