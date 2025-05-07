package pana

import (
	"encoding/json"

	ld "sourcery.dny.nu/longdistance"
	secv1 "sourcery.dny.nu/pana/vocab/w3id/securityv1"
)

// PublicKey is the W3ID Security v1 Public Key.
type PublicKey ld.Node

// NewPublicKey initialises a new PublicKey.
func NewPublicKey() *PublicKey {
	return &PublicKey{
		Properties: make(ld.Properties, 2),
		Type:       []string{secv1.PublicKey},
	}
}

// Build finalises the PublicKey.
func (pk *PublicKey) Build() PublicKey {
	return *pk
}

// GetID returns the ID in [ld.Node].
func (pk *PublicKey) GetID() string {
	return pk.ID
}

// SetID sets the ID in [ld.Node].
func (pk *PublicKey) SetID(id string) *PublicKey {
	pk.ID = id
	return pk
}

// GetType returns the first type on the [ld.Node].
func (pk *PublicKey) GetType() string {
	return pk.Type[0]
}

// SetType sets the type on the [ld.Node].
func (pk *PublicKey) SetType(id string) *PublicKey {
	pk.Type = []string{id}
	return pk
}

// GetOwner returns the ID in [secv1.Owner].
func (pk *PublicKey) GetOwner() string {
	if nodes := (*ld.Node)(pk).GetNodes(secv1.Owner); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetOwner sets the ID in [secv1.Owner].
func (pk *PublicKey) SetOwner(id string) *PublicKey {
	(*ld.Node)(pk).SetNodes(secv1.Owner, ld.Node{ID: id})
	return pk
}

// GetPublicKeyPEM returns the value in [secv1.PublicKeyPem].
func (pk *PublicKey) GetPublicKeyPEM() json.RawMessage {
	if nodes := (*ld.Node)(pk).GetNodes(secv1.PublicKeyPem); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetPublicKeyPEM sets the value in [secv1.PublicKeyPem].
func (pk *PublicKey) SetPublicKeyPEM(v json.RawMessage) *PublicKey {
	(*ld.Node)(pk).SetNodes(secv1.PublicKeyPem, ld.Node{Value: v})
	return pk
}
