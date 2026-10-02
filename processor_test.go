package pana_test

import (
	"errors"
	"strings"
	"testing"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana"
)

func TestValidateContext(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantErr error
	}{
		{
			name: "mastodon create note",
			in: `{
				"@context": [
					"https://www.w3.org/ns/activitystreams",
					{
						"ostatus": "http://ostatus.org#",
						"atomUri": "ostatus:atomUri",
						"inReplyToAtomUri": "ostatus:inReplyToAtomUri",
						"conversation": "ostatus:conversation",
						"sensitive": "as:sensitive",
						"toot": "http://joinmastodon.org/ns#",
						"votersCount": "toot:votersCount"
					},
					"https://w3id.org/security/v1"
				],
				"id": "https://example.com/activity/1",
				"type": "Create"
			}`,
		},
		{
			name: "mastodon update person",
			in: `{
				"@context": [
					"https://www.w3.org/ns/activitystreams",
					"https://w3id.org/security/v1",
					{
						"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
						"toot": "http://joinmastodon.org/ns#",
						"featured": {"@id": "toot:featured", "@type": "@id"},
						"featuredTags": {"@id": "toot:featuredTags", "@type": "@id"},
						"alsoKnownAs": {"@id": "as:alsoKnownAs", "@type": "@id"},
						"movedTo": {"@id": "as:movedTo", "@type": "@id"},
						"schema": "http://schema.org#",
						"PropertyValue": "schema:PropertyValue",
						"value": "schema:value",
						"discoverable": "toot:discoverable",
						"suspended": "toot:suspended",
						"memorial": "toot:memorial",
						"indexable": "toot:indexable",
						"attributionDomains": {"@id": "toot:attributionDomains", "@type": "@id"},
						"focalPoint": {"@container": "@list", "@id": "toot:focalPoint"}
					}
				],
				"id": "https://example.com/activity/2",
				"type": "Update"
			}`,
		},
		{
			name: "redefined activitystreams term",
			in: `{
				"@context": [
					"https://www.w3.org/ns/activitystreams",
					{"content": "as:summary"}
				],
				"id": "https://example.com/activity/3",
				"type": "Create"
			}`,
			wantErr: ld.ErrInvalid,
		},
		{
			name: "redefined security v1 term",
			in: `{
				"@context": [
					"https://www.w3.org/ns/activitystreams",
					"https://w3id.org/security/v1",
					{"publicKeyPem": "sec:owner"}
				],
				"id": "https://example.com/activity/4",
				"type": "Update"
			}`,
			wantErr: ld.ErrInvalid,
		},
	}

	proc := pana.New(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := proc.Unmarshal(t.Context(), strings.NewReader(tt.in), "")
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
