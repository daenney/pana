// Package pana can be used to process [ActivityStreams] messages.
//
// This library is meant to help with implementing fediverse ([ActivityPub])
// clients and servers.
//
// # Usage
//
// An application starts with creating a [Processor] using [NewProcessor]. This
// processor should exist for the lifetime of the application, and not a single
// request.
//
// Incoming messages can then be processed by calling [Processor.Unmarshal].
// This will return an [Activity] that you can inspect and drill into.
//
// Call [Activity.GetType] to figure out what type of activity it is. With that
// information you can then cast the activity to a more specific type which will
// add new getters and setters for the properties that are set on that activity.
// You can use [Properties] to get a set of all properties set on the activity.
// This lets you handle any property you're interested in.
//
// If an activity has an [Object], you can do the same thing. Use
// [Activity.GetObject] to get the object. An object may be a reference, or a
// full object. You can check this with [Object.IsReference]. Once you have a
// full object, you can use [Object.GetType] to determine the type and cast it
// to a more specific type. And you can use [Properties] to determine the
// properties set on the object.
//
// Once you've created an [Object] you wrap it in an [Activity] and call
// [Processor.Marshal]. You can then exchange the resulting JSON using the
// client-to-server or server-to-server APIs.
//
// If you're going to store ActivityStreams messages, ensure you always store
// the original you received. Should an implementation bug in Unmarshal result
// in a problem, you can then always reprocess messages after that's been fixed.
// But if you store a potentially incorrectly unmarshalled document you may not
// be able to return to its original form even if you retain the context.
//
// # JSON-LD
//
// ActivityStreams uses JSON-LD for its JSON serialisatoin format. Pana doesn't
// hide this from you, but aims to provide an API that makes working with this
// simple enough so that you don't have to think about it.
//
// Pana's Unmarshal converts a document to JSON-LD Expanded Document Form. The
// Marshal method does the opposite, converting to JSON-LD Compacted Document
// Form.
//
// In JSON-LD all properties, with the exception of 'id' / '@id', are arrays.
// But if an array only has a single etnry it's typically reduced to its member.
// This distinction is removed in expanded document form, but it makes the
// resulting [Object] a bit verbose to handle.
//
// # API
//
// The API Pana exposes is meant to do the right thing for you in 99.9% of
// cases. Its API is based on what properties are used in practice across the
// fediverse, and what their values are. This is a more restrictive subset of
// what the ActivityStreams specification allows for. Pana does this on purpose
// to guide you towards maximal interoperability with other, potentially JSON-LD
// unaware, implementations.
//
// For any proprety, you'll either have:
//   - GetXXX and SetXXX for single-valued properties. These will accept and
//     return strings or [encoding/json.RawMessage].
//   - GetXXX and AddXXX for multi-valued properties. Get will return an
//     [iter.Seq], and Add will be variadic. Add appends so can be called
//     multiple times too.
//
// Most value properties return [encoding/json.RawMessage]. You can look at the
// JSON-LD context definition to determine what the type of the value should be.
// This will also be part of the vocabulary documentation for the property
// you're retrieving. However, remember that this is not enforced, so even if a
// property is defined as a number it can hold a boolean.
//
// All types in Pana are a [sourcery.dny.nu/longdistance.Node]. This struct type
// has exported fields for every JSON-LD keyword, as well as a catch-all
// Properties field for all other properties. You can directly manipulate any of
// these, but at that point you're responsible for ensuring you create valid
// JSON-LD nodes that will compact to a representation non-JSON-LD aware
// processors will understand. In order to Get, Add or Set properties you need
// to use the IRI a property maps to. This is part of what the JSON-LD context
// defines. The [sourcery.dny.nu/pana/vocab] packages provide the necessary consts.
//
// [ActivityStreams]: https://www.w3.org/TR/activitystreams-core/
// [ActivityPub]: https://www.w3.org/TR/activitypub/
package pana
