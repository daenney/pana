package pana

import (
	"encoding/json"
	"iter"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab/litepub"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// IntransitiveActivity is the ActivityStreams Intransitive Activity type.
//
// It's aliassed to [Object] because that's basically what it is. In practice
// this is rarely if ever used. Polls are modelled as Create Question, even
// though Question is intransitive.
type IntransitiveActivity = Object

// Activity is the ActivityStreams Activity type.
//
// This is a more limited view of [Object] with property getters and setters for
// the properties that are commonly set on the toplevel activity.
type Activity Object

// NewActivity initialises a new activity.
//
// It's initialised with [as.TypeCreate]. Use [Activity.SetType] to override it.
func NewActivity() *Activity {
	return &Activity{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeCreate},
	}
}

// Accept is the Accept activity.
type Accept = Activity

// NewAccept initialises a new Accept activity with [as.TypeAccept].
func NewAccept() *Accept {
	return (*Accept)(NewActivity().SetType(as.TypeAccept))
}

type Add = Activity

// NewAdd initialises a new Add activity with [as.TypeAdd].
func NewAdd() *Add {
	return (*Add)(NewActivity().SetType(as.TypeAdd))
}

type Announce = Activity

// NewAnnounce initialises a new Announce activity with [as.TypeAnnounce].
func NewAnnounce() *Announce {
	return (*Announce)(NewActivity().SetType(as.TypeAnnounce))
}

type Block = Activity

// NewBlock initialises a new Block activity with [as.TypeBlock].
func NewBlock() *Block {
	return (*Block)(NewActivity().SetType(as.TypeBlock))
}

type Create = Activity

// NewCreate initialises a new Create activity with [as.TypeCreate].
func NewCreate() *Create {
	return NewActivity()
}

type Delete = Activity

// NewDelete initialises a new Delete activity with [as.TypeDelete].
func NewDelete() *Delete {
	return (*Delete)(NewActivity().SetType(as.TypeDelete))
}

type EmojiReact = Activity

// NewEmojiReact initialises a new EmojiReact activity with
// [litepub.TypeEmojiReact].
func NewEmojiReact() *EmojiReact {
	return (*EmojiReact)(NewActivity().SetType(litepub.TypeEmojiReact))
}

type Follow = Activity

// NewFollow initialises a new Follow activity with [as.TypeFollow].
func NewFollow() *Follow {
	return (*Follow)(NewActivity().SetType(as.TypeFollow))
}

type Like = Activity

// NewLike initialises a new Like activity with [as.TypeLike].
func NewLike() *Like {
	return (*Like)(NewActivity().SetType(as.TypeLike))
}

type Move = Activity

// NewMove initialises a new Move activity with [as.TypeMove].
func NewMove() *Move {
	return (*Move)(NewActivity().SetType(as.TypeMove))
}

type Remove = Activity

// NewRemove initialises a new Remove activity with [as.TypeRemove].
func NewRemove() *Remove {
	return (*Remove)(NewActivity().SetType(as.TypeRemove))
}

type Undo = Activity

// NewUndo initialises a new Undo activity with [as.TypeUndo].
func NewUndo() *Undo {
	return (*Undo)(NewActivity().SetType(as.TypeUndo))
}

type Update = Activity

// NewUpdate initialises a new Update activity with [as.TypeUpdate].
func NewUpdate() *Update {
	return (*Update)(NewActivity().SetType(as.TypeUpdate))
}

// Build finalises the Activity.
func (a *Activity) Build() Activity {
	return *a
}

// IsIntransitive returns true if the activity has no object.
func (a *Activity) IsIntransitive() bool {
	return !Has(a, as.Object)
}

// See [Object.GetID].
func (a *Activity) GetID() string {
	return (*Object)(a).GetID()
}

// See [Object.SetID].
func (a *Activity) SetID(id string) *Activity {
	(*Object)(a).SetID(id)
	return a
}

// See [Object.GetType].
func (a *Activity) GetType() string {
	return (*Object)(a).GetType()
}

// See [Object.SetType].
func (a *Activity) SetType(typ string) *Activity {
	(*Object)(a).SetType(typ)
	return a
}

// GetActor returns the actor IDs from [as.Actor].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-actor.
func (a *Activity) GetActor() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, n := range (*ld.Node)(a).GetNodes(as.Actor) {
			if !yield(n.ID) {
				return
			}
		}
	}
}

// AddActor appends actor IDs to [as.Actor].
func (a *Activity) AddActor(ids ...string) *Activity {
	(*ld.Node)(a).AddNodes(as.Actor, toReference(ids...)...)
	return a
}

// GetInstrument returns the instrument in [as.Instrument].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-instrument.
func (a *Activity) GetInstrument() *Instrument {
	if nodes := (*ld.Node)(a).GetNodes(as.Instrument); len(nodes) == 1 {
		return (*Instrument)(&nodes[0])
	}

	return nil
}

// SetInstrument sets the instrument in [as.Instrument].
func (a *Activity) SetInstrument(in Instrument) *Activity {
	(*ld.Node)(a).SetNodes(as.Instrument, ld.Node(in))
	return a
}

// GetObject returns the object in [as.Object].
//
// This returns [as.Any] because it can be of many different types. If the
// [Any.GetType] doesn't match any known type you can cast it to [Object].
//
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-object.
func (a *Activity) GetObject() *Any {
	if nodes := (*ld.Node)(a).GetNodes(as.Object); len(nodes) == 1 {
		return (*Any)(&nodes[0])
	}

	return nil
}

// SetObject sets the object in [as.Object].
//
// It is possible to have more than one object. However, except for JSON-LD
// aware processors, nobody understands this. If you want to send multiple
// objects, send multiple activities instead.
func (a *Activity) SetObject(obj Any) *Activity {
	(*ld.Node)(a).SetNodes(as.Object, ld.Node(obj))
	return a
}

// GetTarget returns the ID in [as.Target].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-target.
func (a *Activity) GetTarget() string {
	if nodes := (*ld.Node)(a).GetNodes(as.Target); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetTarget sets the ID in [as.Target].
func (a *Activity) SetTarget(id string) *Activity {
	(*ld.Node)(a).SetNodes(as.Target, ld.Node{ID: id})
	return a
}

// See [Object.GetCc].
func (a *Activity) GetCc() iter.Seq[string] {
	return (*Object)(a).GetCc()
}

// See [Object.AddCc].
func (a *Activity) AddCc(ids ...string) *Activity {
	(*Object)(a).AddCc(ids...)
	return a
}

// See [Object.GetTo].
func (a *Activity) GetTo() iter.Seq[string] {
	return (*Object)(a).GetTo()
}

// See [Object.AddTo].
func (a *Activity) AddTo(ids ...string) *Activity {
	(*Object)(a).AddTo(ids...)
	return a
}

// See [Object.GetPublished].
func (a *Activity) GetPublished() json.RawMessage {
	return (*Object)(a).GetPublished()
}

// See [Object.SetPublished].
func (a *Activity) SetPublished(v json.RawMessage) *Activity {
	(*Object)(a).SetPublished(v)
	return a
}
