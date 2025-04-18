package pana

import (
	"encoding/json"

	ld "code.dny.dev/longdistance"
	secv1 "code.dny.dev/pana/vocab/w3id/securityv1"
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

// See [Object.GetID].
func (pk *PublicKey) GetID() string {
	return pk.ID
}

// See [Object.SetID].
func (pk *PublicKey) SetID(id string) *PublicKey {
	pk.ID = id
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
