package pana_test

import (
	"encoding/json"
	"fmt"
	"iter"
	"log/slog"
	"maps"
	"slices"

	"code.dny.dev/pana"
	"code.dny.dev/pana/vocab/mastodon"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
)

func ExampleProcessor_Marshal() {
	proc := pana.NewProcessor(slog.New(slog.DiscardHandler))

	activity := pana.NewActivity().
		SetID("https://example.com/id/1").
		AddTo(as.PublicCollection).
		SetObject(
			pana.NewNote().
				AddContent(
					pana.NewLocalised().
						SetValue([]byte(`"We spell color wrong."`)).
						SetLanguage("en-us").
						Build(),
				).
				AddContent(
					pana.NewLocalised().
						SetValue([]byte(`"Omelette du fromage."`)).
						SetLanguage("fr-FR").
						Build(),
					pana.NewLocalised().
						SetValue([]byte(`"Ona li toki."`)).
						Build(),
				).
				Build()).
		SetType(as.TypeCreate).
		Build()

	compacted, err := proc.Marshal(
		json.RawMessage(`{"@context":"https://www.w3.org/ns/activitystreams"}`),
		activity,
	)

	if err != nil {
		panic(err)
	}

	var res any
	_ = json.Unmarshal(compacted, &res)
	ind, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(ind))
	// Output:
	// {
	//   "@context": "https://www.w3.org/ns/activitystreams",
	//   "id": "https://example.com/id/1",
	//   "object": {
	//     "content": "Ona li toki.",
	//     "contentMap": {
	//       "en-us": "We spell color wrong.",
	//       "fr-FR": "Omelette du fromage."
	//     },
	//     "type": "Note"
	//   },
	//   "to": "https://www.w3.org/ns/activitystreams#Public",
	//   "type": "Create"
	// }
}

func ExampleProcessor_Unmarshal() {
	// small helper for stable iteration order of properties
	var listProperties = func(in map[string]struct{}) []string {
		s := slices.Collect(maps.Keys(in))
		slices.Sort(s)
		return s
	}

	var getSingle = func(in iter.Seq[*pana.Localised]) string {
		for val := range in {
			return string(val.Value)
		}
		return ""
	}

	msg := []byte(`{
	"@context": [
        "https://www.w3.org/ns/activitystreams",
		{
			"Emoji": "toot:Emoji",
			"Hashtag": "as:Hashtag",
			"atomUri": "ostatus:atomUri",
			"blurhash": "toot:blurhash",
			"focalPoint": {
				"@id": "toot:focalPoint",
				"@container": "@list"
			},
			"ostatus": "http://ostatus.org#",
			"toot": "http://joinmastodon.org/ns#",
			"sensitive": "as:sensitive"
		}
    ],
    "actor": "https://example.com/actor/wcgkaaorsn",
    "cc": "https://www.w3.org/ns/activitystreams#Public",
    "id": "https://example.com/id/bikfmutagl",
    "object": {
		"atomUri": "https://example.com/atomUri/qbfujvxbeo",
        "attachment": {
			"blurhash": "pkbjlvkjjb",
			"focalPoint": [
				0.7,
				-0.7
			],
			"height": 600,
			"mediaType": "image/jpeg",
			"name": "banana",
			"type": "Document",
			"url": "https://example.com/url/odybnszymk",
			"width": 800
        },
        "attributedTo": "https://example.com/attributedTo/ytoxvxmalh",
        "content": "eating a banana",
        "id": "https://example.com/id/bfcqgpkwjv",
        "published": "2020-10-01T16:00:00Z",
        "replies": {
            "first": {
                "id": "https://example.com/id/htpqbwrroa",
                "next": "https://example.com/next/dsqoqssppw",
                "partOf": "https://example.com/partOf/bjkaungqre",
                "type": "CollectionPage"
            },
            "id": "https://example.com/id/edjlqpbspm",
            "type": "Collection"
        },
        "sensitive": true,
        "summary": "food, eye contact",
        "tag": [
            {
                "icon": {
                    "mediaType": "image/jpeg",
                    "type": "Image",
                    "url": "https://example.com/url/kkurwzofjc"
                },
                "id": "https://example.com/id/xadgmrnzdk",
                "name": "bananadance",
                "type": "Emoji",
                "updated": "2020-10-01T16:00:00Z"
            },
            {
                "href": "https://example.com/href/qzqgetcjfv",
                "name": "#food",
                "type": "Hashtag"
            }
        ],
        "to": "https://www.w3.org/ns/activitystreams#Public",
        "type": "Note",
        "url": "https://example.com/url/lpsbcbqdnm"
    },
    "published": "2020-10-01T16:00:00Z",
    "to": "https://www.w3.org/ns/activitystreams#Public",
    "type": "Create"
}`)
	proc := pana.NewProcessor(slog.New(slog.DiscardHandler))
	activity, err := proc.Unmarshal(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("#----- Activity -----#")
	// what type of Activity did we just get?
	fmt.Println("type:", activity.GetType())

	// what properties are set on the Activity?
	for _, property := range listProperties(pana.Properties(activity)) {
		fmt.Println("property:", property)
	}

	fmt.Println("\n#----- Object -----#")
	// what's our object?
	obj := activity.GetObject()

	// do we have a real object or a reference?
	fmt.Println("object is a reference:", pana.IsReference(obj))

	// what's the object type?
	fmt.Println("object type:", obj.GetType())

	// what are the properties on the object?
	for _, property := range listProperties(pana.Properties(obj)) {
		fmt.Println("property:", property)
	}

	// time to convert it
	note := (*pana.Note)(obj) // Note is a more limited pana.Object with only the getters and setters for microblog posts.

	fmt.Println("\n#----- Object attachments -----#")
	for atch := range note.GetAttachment() {
		switch atch.GetType() {
		case as.TypeDocument:
			doc := (*pana.Document)(atch)
			fmt.Println("height:", string(doc.GetHeight()))
			fmt.Println("width:", string(doc.GetWidth()))
			fmt.Println("URL:", doc.GetURL())
		default:
			panic("unknown attachment type")
		}
	}

	fmt.Println("\n#----- Object tags -----#")
	// lets get all the tags
	for tag := range note.GetTag() {
		switch tag.GetType() {
		case as.TypeHashtag:
			ht := (*pana.Hashtag)(tag)
			fmt.Println("hashtag name:", getSingle(ht.GetName()))
			fmt.Println("hashtag IRI:", ht.GetHref())
		case mastodon.TypeEmoji:
			emo := (*pana.Emoji)(tag)
			fmt.Println("emoji name:", getSingle(emo.GetName()))
			fmt.Println("emoj URL:", emo.GetIcon().GetURL())
		default:
			panic("unsupported tag type")
		}
	}

	// Output:
	// #----- Activity -----#
	// type: https://www.w3.org/ns/activitystreams#Create
	// property: https://www.w3.org/ns/activitystreams#actor
	// property: https://www.w3.org/ns/activitystreams#cc
	// property: https://www.w3.org/ns/activitystreams#object
	// property: https://www.w3.org/ns/activitystreams#published
	// property: https://www.w3.org/ns/activitystreams#to
	// property: id
	// property: type
	//
	// #----- Object -----#
	// object is a reference: false
	// object type: https://www.w3.org/ns/activitystreams#Note
	// property: http://ostatus.org#atomUri
	// property: https://www.w3.org/ns/activitystreams#attachment
	// property: https://www.w3.org/ns/activitystreams#attributedTo
	// property: https://www.w3.org/ns/activitystreams#content
	// property: https://www.w3.org/ns/activitystreams#published
	// property: https://www.w3.org/ns/activitystreams#replies
	// property: https://www.w3.org/ns/activitystreams#sensitive
	// property: https://www.w3.org/ns/activitystreams#summary
	// property: https://www.w3.org/ns/activitystreams#tag
	// property: https://www.w3.org/ns/activitystreams#to
	// property: https://www.w3.org/ns/activitystreams#url
	// property: id
	// property: type
	//
	// #----- Object attachments -----#
	// height: 600
	// width: 800
	// URL: https://example.com/url/odybnszymk
	//
	// #----- Object tags -----#
	// emoji name: "bananadance"
	// emoj URL: https://example.com/url/kkurwzofjc
	// hashtag name: "#food"
	// hashtag IRI: https://example.com/href/qzqgetcjfv
}
