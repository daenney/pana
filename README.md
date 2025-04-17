# Pana

A Go library for building on the Fediverse.

Pana provides everything you need to process and create [ActivityStreams][as] messages. These are the kinds of messages that are exchanged between clients and servers on the Fediverse.

[as]: https://www.w3.org/TR/activitystreams-core/

The API Pana provides is tailored towards social media style applications. The underlying library, [longdistance][ld], can be used for anything that use JSON-LD.

[ld]: https://codeberg.org/daenney/longdistance

## Usage

The [Go documentation][godoc] includes a [few examples][godocex].

[godoc]: https://pkg.go.dev/code.dny.dev/pana
[godocex]: https://pkg.go.dev/code.dny.dev/pana#pkg-examples

## Contributing

PRs are very welcome for:
* Completing current ActivityStreams types.
* Fixes to existing types.
* Adding new object types that are in use on the Fediverse.
* Documentation improvements.

If you run into a bug, feel free to open up an issue.

## License

This library is licensed under the Mozilla Public License Version 2.0.
