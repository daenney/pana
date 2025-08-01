package pana

import (
	"encoding/json"
	"iter"
	"time"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/mastodon"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
)

// Question is the ActivityStreams Question object.
type Question IntransitiveActivity

// NewQuestion initialises a new Question.
func NewQuestion() *Question {
	return &Question{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeQuestion},
	}
}

// Build finalises the Question.
func (q *Question) Build() Question {
	return *q
}

// IsMultipleChoice checks if this is a multiple-choice question.
func (q *Question) IsMultipleChoice() bool {
	if q == nil {
		return false
	}

	return Has((*Any)(q), as.AnyOf)
}

// See [Object.GetType].
func (q *Question) GetType() string {
	return (*Object)(q).GetType()
}

// SetType sets the type to [as.TypeQuestion].
func (q *Question) SetType() {
	q.Type = []string{as.TypeQuestion}
}

// GetVotersCount returns the value in [mastodon.VotersCount].
//
// See https://docs.joinmastodon.org/spec/activitypub/#toot.
func (q *Question) GetVotersCount() json.RawMessage {
	if nodes := (*ld.Node)(q).GetNodes(mastodon.VotersCount); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetVotersCount sets the number in [mastodon.VotersCount].
func (q *Question) SetVotersCount(v uint64) *Question {
	data, _ := json.Marshal(v)
	(*ld.Node)(q).SetNodes(mastodon.VotersCount, ld.Node{Value: data})
	return q
}

// SetVotersCountRaw sets the value in [mastodon.VotersCount].
func (q *Question) SetVotersCountRaw(v json.RawMessage) *Question {
	(*ld.Node)(q).SetNodes(mastodon.VotersCount, ld.Node{Value: v})
	return q
}

// GetEndTime returns the value in [as.EndTime].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-endtime.
func (q *Question) GetEndTime() json.RawMessage {
	if nodes := (*ld.Node)(q).GetNodes(as.EndTime); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetEndTime sets the [time.Time] in [as.EndTime].
func (q *Question) SetEndTime(v time.Time) *Question {
	data, _ := json.Marshal(v.Format(time.RFC3339))
	(*ld.Node)(q).SetNodes(as.EndTime, ld.Node{Value: data, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// SetEndTimeRaw sets the value in [as.EndTime].
func (q *Question) SetEndTimeRaw(v json.RawMessage) *Question {
	(*ld.Node)(q).SetNodes(as.EndTime, ld.Node{Value: v, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// See [Object.GetSensitive].
func (q *Question) GetSensitive() json.RawMessage {
	return (*Object)(q).GetSensitive()
}

// See [Object.SetSensitive].
func (q *Question) SetSensitive(v bool) *Question {
	(*Object)(q).SetSensitive(v)
	return q
}

// See [Object.SetSensitiveRaw].
func (q *Question) SetSensitiveRaw(v json.RawMessage) *Question {
	(*Object)(q).SetSensitiveRaw(v)
	return q
}

// See [Object.GetUpdated].
func (q *Question) GetUpdated() json.RawMessage {
	return (*Object)(q).GetUpdated()
}

// See [Object.SetUpdated].
func (q *Question) SetUpdated(v time.Time) *Question {
	(*Object)(q).SetUpdated(v)
	return q
}

// See [Object.SetUpdatedRaw].
func (q *Question) SetUpdatedRaw(v json.RawMessage) *Question {
	(*Object)(q).SetUpdatedRaw(v)
	return q
}

// GetAnyOf gets the [Choice] in [as.AnyOf].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-oneof.
func (q *Question) GetAnyOf() iter.Seq[Choice] {
	return func(yield func(Choice) bool) {
		for _, n := range (*ld.Node)(q).GetNodes(as.AnyOf) {
			if !yield(Choice(n)) {
				return
			}
		}
	}
}

// AddAnyOf appends [Choice] to [as.AnyOf].
func (q *Question) AddAnyOf(chs ...Choice) *Question {
	(*ld.Node)(q).AddNodes(as.AnyOf, toLDNodes(chs...)...)
	return q
}

// GetOneOf gets the [Choice] in [as.OneOf].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-oneof.
func (q *Question) GetOneOf() iter.Seq[Choice] {
	return func(yield func(Choice) bool) {
		for _, n := range (*ld.Node)(q).GetNodes(as.OneOf) {
			if !yield(Choice(n)) {
				return
			}
		}
	}
}

// AddOneOf appends [Choice] to [as.OneOf].
func (q *Question) AddOneOf(chs ...Choice) *Question {
	(*ld.Node)(q).AddNodes(as.OneOf, toLDNodes(chs...)...)
	return q
}

// GetClosed returns the value in [as.Closed].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-closed.
func (q *Question) GetClosed() json.RawMessage {
	if nodes := (*ld.Node)(q).GetNodes(as.Closed); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetClosed sets the [time.Time] in [as.Closed].
func (q *Question) SetClosed(v time.Time) *Question {
	data, _ := json.Marshal(v.Format(time.RFC3339))
	(*ld.Node)(q).SetNodes(as.Closed, ld.Node{Value: data, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// SetClosedRawsets the value in [as.Closed].
func (q *Question) SetClosedRaw(v json.RawMessage) *Question {
	(*ld.Node)(q).SetNodes(as.Closed, ld.Node{Value: v, Type: []string{xmlschema.TypeDateTime}})
	return q
}

// See [Object.GetName].
func (q *Question) GetName() iter.Seq[*Localised] {
	return (*Object)(q).GetName()
}

// See [Object.AddName].
func (q *Question) AddName(ls ...Localised) *Question {
	(*Object)(q).AddName(ls...)
	return q
}

// Choice represents a choice in a poll. It is a much more limited [Note].
type Choice Note

// NewChoice initialises a new Choice.
func NewChoice() *Choice {
	return &Choice{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeNote},
	}
}

// Build finalises the Choice.
func (c *Choice) Build() Choice {
	return *c
}

// See [Object.GetName].
func (c *Choice) GetName() iter.Seq[*Localised] {
	return (*Object)(c).GetName()
}

// See [Object.AddName].
func (c *Choice) AddName(ls ...Localised) *Choice {
	(*Object)(c).AddName(ls...)
	return c
}

// See [Object.GetReplies].
func (c *Choice) GetReplies() *Collection {
	return (*Object)(c).GetReplies()
}

// See [Object.SetReplies].
func (c *Choice) SetReplies(cl Collection) *Choice {
	(*Object)(c).SetReplies(cl)
	return c
}
