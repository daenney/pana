// Package gotosocial contains terms for the GoToSocial namespace.
package gotosocial

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://gotosocial.org/ns"

// Namespace is the IRI prefix used for terms defined in this namespace.
const Namespace = IRI + "#"

const (
	// Always is an IRI, either as a string or as an object with an
	// id property.
	Always = Namespace + "always"
	// ApprovalRequired is an IRI, either as a string or as an object with an
	// id property.
	ApprovalRequired = Namespace + "approvalRequired"
	// ApprovedBy is an IRI, either as a string or as an object with an
	// id property.
	ApprovedBy = Namespace + "approvedBy"
	// CanAnnounce is an IRI, either as a string or as an object with an
	// id property.
	CanAnnounce = Namespace + "canAnnounce"
	// CanLike is an IRI, either as a string or as an object with an
	// id property.
	CanLike = Namespace + "canLike"
	// CanReply is an IRI, either as a string or as an object with an
	// id property.
	CanReply = Namespace + "canReply"
	// InteractionPolicy is an IRI, either as a string or as an object with an
	// id property.
	InteractionPolicy = Namespace + "interactionPolicy"
	// TypeAnnounceApproval is a possible value for the type property.
	TypeAnnounceApproval = Namespace + "AnnounceApproval"
	// TypeLikeApproval is a possible value for the type property.
	TypeLikeApproval = Namespace + "LikeApproval"
	// TypeReplyApproval is a possible value for the type property.
	TypeReplyApproval = Namespace + "ReplyApproval"
)
