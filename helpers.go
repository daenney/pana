package pana

import (
	"slices"

	ld "sourcery.dny.nu/longdistance"
)

type node interface {
	ld.Node | Object | Article | Activity | Note | Collection | CollectionPage |
		Tombstone | Instrument | Document | Audio | Icon | Link | LinkTag |
		PublicKey | Actor | Any | Page | Place | Emoji | Endpoints | Localised |
		Question | Choice | Event | PropertyValue
}

func toReference(ids ...string) []ld.Node {
	if len(ids) == 0 {
		return nil
	}

	res := make([]ld.Node, 0, len(ids))
	for _, id := range ids {
		res = append(res, ld.Node{ID: id})
	}

	return res
}

func addNodes[T node](n *ld.Node, property string, ins []T) {
	if len(ins) == 0 {
		return
	}

	res := slices.Grow(n.Properties[property], len(ins))
	for _, in := range ins {
		res = append(res, ld.Node(in))
	}

	n.Properties[property] = res
}
