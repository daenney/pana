// Package secv1 contains terms for the W3ID Security namespace.
package secv1

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://w3id.org/security/v1"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/security#"

const (
	// AuthenticationTag is a string.
	AuthenticationTag = Namespace + "authenticationTag"
	// CanonicalizationAlgorithm is a string.
	CanonicalizationAlgorithm = Namespace + "canonicalizationAlgorithm"
	// CipherAlgorithm is a string.
	CipherAlgorithm = Namespace + "cipherAlgorithm"
	// CipherData is a string.
	CipherData = Namespace + "cipherData"
	// CipherKey is a string.
	CipherKey = Namespace + "cipherKey"
	// DigestAlgorithm is a string.
	DigestAlgorithm = Namespace + "digestAlgorithm"
	// DigestValue is a string.
	DigestValue = Namespace + "digestValue"
	// Domain is a string.
	Domain = Namespace + "domain"
	// EncryptionKey is a string.
	EncryptionKey = Namespace + "encryptionKey"
	// Expiration is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Expiration = Namespace + "expiration"
	// Expires is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Expires = Namespace + "expiration"
	// InitializationVector is a string.
	InitializationVector = Namespace + "initializationVector"
	// IterationCount is a string.
	IterationCount = Namespace + "iterationCount"
	// Nonce is a string.
	Nonce = Namespace + "nonce"
	// NormalizationAlgorithm is a string.
	NormalizationAlgorithm = Namespace + "normalizationAlgorithm"
	// Owner is an IRI, either as a string or as an object with an
	// id property.
	Owner = Namespace + "owner"
	// Password is a string.
	Password = Namespace + "password"
	// PrivateKey is an IRI, either as a string or as an object with an
	// id property.
	PrivateKey = Namespace + "privateKey"
	// PrivateKeyPem is a string.
	PrivateKeyPem = Namespace + "privateKeyPem"
	// PublicKey is an IRI, either as a string or as an object with an
	// id property.
	PublicKey = Namespace + "publicKey"
	// PublicKeyBase58 is a string.
	PublicKeyBase58 = Namespace + "publicKeyBase58"
	// PublicKeyPem is a string.
	PublicKeyPem = Namespace + "publicKeyPem"
	// PublicKeyService is an IRI, either as a string or as an object with an
	// id property.
	PublicKeyService = Namespace + "publicKeyService"
	// PublicKeyWif is a string.
	PublicKeyWif = Namespace + "publicKeyWif"
	// Revoked is an xml:dateTime, equivalent to a time.Date in RFC3339Nano.
	Revoked = Namespace + "revoked"
	// Salt is a string.
	Salt = Namespace + "salt"
	// Signature is a string.
	Signature = Namespace + "signature"
	// SignatureAlgorithm is a string.
	SignatureAlgorithm = Namespace + "signingAlgorithm"
	// SignatureValue is a string.
	SignatureValue = Namespace + "signatureValue"
	// TypeCryptographicKey is a possible value for the type property.
	TypeCryptographicKey = Namespace + "Key"
	// TypeEcdsaKoblitzSignature2016 is a possible value for the type property.
	TypeEcdsaKoblitzSignature2016 = Namespace + "EcdsaKoblitzSignature2016"
	// TypeEd25519Signature2018 is a possible value for the type property.
	TypeEd25519Signature2018 = Namespace + "Ed25519Signature2018"
	// TypeEncryptedMessage is a possible value for the type property.
	TypeEncryptedMessage = Namespace + "EncryptedMessage"
	// TypeGraphSignature2012 is a possible value for the type property.
	TypeGraphSignature2012 = Namespace + "GraphSignature2012"
	// TypeLinkedDataSignature2015 is a possible value for the type property.
	TypeLinkedDataSignature2015 = Namespace + "LinkedDataSignature2015"
	// TypeLinkedDataSignature2016 is a possible value for the type property.
	TypeLinkedDataSignature2016 = Namespace + "LinkedDataSignature2016"
)
