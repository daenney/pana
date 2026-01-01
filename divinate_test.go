package pana_test

import (
	"encoding/json"
	"reflect"
	"testing"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana"
	"sourcery.dny.nu/pana/vocab/mastodon"
	"sourcery.dny.nu/pana/vocab/schema"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/ldp"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
	secv1 "sourcery.dny.nu/pana/vocab/w3id/securityv1"
)

func TestDivinate(t *testing.T) {
	tests := []struct {
		name string
		in   ld.Node
		want any
	}{
		{
			name: "empty document",
			in:   ld.Node{},
			want: as.IRI,
		},
		{
			name: "core activitystreams properties",
			in: ld.Node{
				Properties: ld.Properties{as.SharedInbox: []ld.Node{
					{ID: "https://example.com"},
				}},
				Type: []string{as.TypeAccept},
			},
			want: as.IRI,
		},
		{
			name: "core activitystreams property and non-core Type",
			in: ld.Node{
				Properties: ld.Properties{as.SharedInbox: []ld.Node{
					{ID: "https://example.com"},
				}},
				Type: []string{mastodon.TypeEmoji},
			},
			want: []any{
				as.IRI,
				map[string]any{
					mastodon.Prefix: mastodon.Namespace,
					mastodon.Term(
						mastodon.TypeEmoji,
					): mastodon.CompactIRI(mastodon.TypeEmoji),
				},
			},
		},
		{
			name: "actor profile",
			in: ld.Node{
				ID:   "https://example.com/id",
				Type: []string{as.TypePerson},
				Properties: ld.Properties{
					as.Following:                 []ld.Node{{ID: "https://example.com/following"}},
					as.Followers:                 []ld.Node{{ID: "https://example.com/followers"}},
					ldp.Inbox:                    []ld.Node{{ID: "https://example.com/indox"}},
					as.Outbox:                    []ld.Node{{ID: "https://example.com/outbox"}},
					mastodon.Featured:            []ld.Node{{ID: "https://example.com/featured"}},
					mastodon.FeaturedTags:        []ld.Node{{ID: "https://example.com/featuredTags"}},
					as.PreferredUsername:         []ld.Node{{Value: []byte(`"username"`)}},
					as.Name:                      []ld.Node{{Value: []byte(`"name"`)}},
					as.Summary:                   []ld.Node{{Value: []byte(`"summary"`)}},
					as.URL:                       []ld.Node{{ID: "https://example.com/url"}},
					as.ManuallyApprovesFollowers: []ld.Node{{Value: []byte(`true`)}},
					mastodon.Discoverable:        []ld.Node{{Value: []byte(`true`)}},
					mastodon.Indexable:           []ld.Node{{Value: []byte(`true`)}},
					as.Published:                 []ld.Node{{Value: []byte(`"2010-01-01T00:00:00Z"`)}},
					as.Sensitive:                 []ld.Node{{Value: []byte(`true`)}},
					mastodon.Memorial:            []ld.Node{{Value: []byte(`true`)}},
					mastodon.AttributionDomains: []ld.Node{
						{ID: "https://example.com/domain1"},
						{ID: "https://example.com/domain2"},
					},
					secv1.PublicKey: []ld.Node{
						{
							ID: "https://example.com/publicKey",
							Properties: ld.Properties{
								secv1.Owner:        []ld.Node{{ID: "https://example.com/owner"}},
								secv1.PublicKeyPem: []ld.Node{{Value: []byte(`"pem"`)}},
							},
						},
					},
					as.Attachment: []ld.Node{
						{
							Type: []string{schema.TypePropertyValue},
							Properties: ld.Properties{
								as.Name:      []ld.Node{{Value: []byte(`"name"`)}},
								schema.Value: []ld.Node{{Value: []byte(`"value"`)}},
							},
						},
					},
					as.Endpoints: []ld.Node{
						{
							Properties: ld.Properties{
								as.SharedInbox: []ld.Node{
									{ID: "https://example.com/sharedIndox"},
								},
							},
						},
					},
					as.Icon: []ld.Node{
						{
							Type: []string{as.TypeImage},
							Properties: ld.Properties{
								as.MediaType: []ld.Node{{Value: []byte(`"mediatype"`)}},
								as.URL:       []ld.Node{{ID: "https://example.com/url"}},
							},
						},
					},
					as.Image: []ld.Node{
						{
							Type: []string{as.TypeImage},
							Properties: ld.Properties{
								as.MediaType: []ld.Node{{Value: []byte(`"mediatype"`)}},
								as.URL:       []ld.Node{{ID: "https://example.com/url"}},
							},
						},
					},
				},
			},
			want: []any{
				as.IRI,
				secv1.IRI,
				map[string]any{
					as.Term(
						as.ManuallyApprovesFollowers,
					): as.CompactIRI(as.ManuallyApprovesFollowers),
					as.Term(as.Sensitive): map[string]any{
						ld.KeywordID:   as.CompactIRI(as.Sensitive),
						ld.KeywordType: xmlschema.CompactIRI(xmlschema.TypeBoolean),
					},
					schema.Prefix: schema.Namespace,
					schema.Term(
						schema.TypePropertyValue,
					): schema.CompactIRI(schema.TypePropertyValue),
					schema.Term(
						schema.Value,
					): schema.CompactIRI(schema.Value),
					mastodon.Prefix: mastodon.Namespace,
					mastodon.Term(
						mastodon.Featured,
					): map[string]any{
						ld.KeywordID:   mastodon.CompactIRI(mastodon.Featured),
						ld.KeywordType: ld.KeywordID,
					},
					mastodon.Term(
						mastodon.FeaturedTags,
					): map[string]any{
						ld.KeywordID:   mastodon.CompactIRI(mastodon.FeaturedTags),
						ld.KeywordType: ld.KeywordID,
					},
					mastodon.Term(
						mastodon.AttributionDomains,
					): map[string]any{
						ld.KeywordID:   mastodon.CompactIRI(mastodon.AttributionDomains),
						ld.KeywordType: ld.KeywordID,
					},
					mastodon.Term(
						mastodon.Discoverable,
					): mastodon.CompactIRI(mastodon.Discoverable),
					mastodon.Term(
						mastodon.Indexable,
					): mastodon.CompactIRI(mastodon.Indexable),
					mastodon.Term(
						mastodon.Memorial,
					): mastodon.CompactIRI(mastodon.Memorial),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := pana.Divinate(tt.in)
			if err != nil {
				t.Fatalf("divination must not fail, got: %q", err.Error())
			}

			var res any
			if err := json.Unmarshal(got, &res); err != nil {
				t.Fatalf("context must not fail to unmarshal, got: %q", err.Error())
			}

			if !reflect.DeepEqual(tt.want, res) {
				t.Errorf("wanted: %#v, got: %#v", tt.want, res)
			}
		})
	}
}
