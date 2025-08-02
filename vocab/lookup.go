package vocab

import (
	"net/url"
	"path"

	dcterms "sourcery.dny.nu/pana/vocab/dublincore/terms"
	"sourcery.dny.nu/pana/vocab/fedibird"
	"sourcery.dny.nu/pana/vocab/geojson"
	gts "sourcery.dny.nu/pana/vocab/gotosocial"
	"sourcery.dny.nu/pana/vocab/litepub"
	"sourcery.dny.nu/pana/vocab/mastodon"
	"sourcery.dny.nu/pana/vocab/misskey"
	"sourcery.dny.nu/pana/vocab/ostatus"
	"sourcery.dny.nu/pana/vocab/schema"
	"sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/ldp"
	rdfs "sourcery.dny.nu/pana/vocab/w3/rdf/schema"
	"sourcery.dny.nu/pana/vocab/w3/vcard"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
	credv1 "sourcery.dny.nu/pana/vocab/w3id/credentialsv1"
	"sourcery.dny.nu/pana/vocab/w3id/identityv1"
	"sourcery.dny.nu/pana/vocab/w3id/payswarmv1"
	permsv1 "sourcery.dny.nu/pana/vocab/w3id/permissionsv1"
	secv1 "sourcery.dny.nu/pana/vocab/w3id/securityv1"
)

// Lookup returns the term definition for an IRI.
//
// It returns the Namespace IRI if the term is defined in that namespace.
// Otherwise it returns a term definition to be added directly to
// the compaction context.
func Lookup(iri string) any {
	u, _ := url.Parse(iri)

	prefix := ""

	if u.Fragment != "" {
		prefix = (&url.URL{Scheme: u.Scheme, Host: u.Host, Path: u.Path}).String()
		prefix += "#"
	} else {
		d, _ := path.Split(u.Path)
		prefix = (&url.URL{Scheme: u.Scheme, Host: u.Host, Path: d}).String()
	}

	switch prefix {
	case activitystreams.Namespace:
		return activitystreams.TermDefForIRI(iri)
	case credv1.Namespace:
		return credv1.TermDefForIRI(iri)
	case dcterms.Namespace:
		return dcterms.TermDefForIRI(iri)
	case fedibird.Namespace:
		return fedibird.TermDefForIRI(iri)
	case geojson.Namespace:
		return geojson.TermDefForIRI(iri)
	case gts.Namespace:
		return gts.TermDefForIRI(iri)
	case identityv1.Namespace:
		return identityv1.TermDefForIRI(iri)
	case ldp.Namespace:
		return ldp.TermDefForIRI(iri)
	case litepub.Namespace:
		return litepub.TermDefForIRI(iri)
	case mastodon.Namespace:
		return mastodon.TermDefForIRI(iri)
	case misskey.Namespace:
		return misskey.TermDefForIRI(iri)
	case ostatus.Namespace:
		return ostatus.TermDefForIRI(iri)
	case payswarmv1.Namespace:
		return payswarmv1.TermDefForIRI(iri)
	case permsv1.Namespace:
		return permsv1.TermDefForIRI(iri)
	case rdfs.Namespace:
		return rdfs.TermDefForIRI(iri)
	case schema.Namespace:
		return schema.TermDefForIRI(iri)
	case secv1.Namespace:
		return secv1.TermDefForIRI(iri)
	case vcard.Namespace:
		return vcard.TermDefForIRI(iri)
	case xmlschema.Namespace:
		return xmlschema.TermDefForIRI(iri)
	default:
		return nil
	}
}
