// Package payswarmv1 contains terms for the W3ID Payswarm namespace.
package payswarmv1

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = "https://w3id.org/payswarm#"

const (
	// PaymentProcessor is a string.
	PaymentProcessor = Namespace + "#processor"
	// Preferences is an IRI, either as a string or as an object with an id
	// property.
	Preferences = Namespace + "#preferences"
)
