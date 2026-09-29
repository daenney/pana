package pana

import (
	"iter"
	"slices"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
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

func getCollectionItem(n *ld.Node, ordered bool) iter.Seq[*Any] {
	return func(yield func(*Any) bool) {
		nodes := n.GetNodes(as.Items)
		if ordered && len(nodes) > 0 {
			nodes = nodes[0].List
		}

		for _, n := range nodes {
			if !yield((*Any)(&n)) {
				return
			}
		}
	}
}

func addCollectionItem[T node](n *ld.Node, ordered bool, items []T) {
	if len(items) == 0 {
		return
	}

	if !ordered {
		addNodes(n, as.Items, items)
		return
	}

	nodes := n.Properties[as.Items]
	if len(nodes) == 0 {
		nodes = []ld.Node{{List: make([]ld.Node, 0, len(items))}}
	}

	list := slices.Grow(nodes[0].List, len(items))
	for _, item := range items {
		list = append(list, ld.Node(item))
	}

	nodes[0].List = list
	n.Properties[as.Items] = nodes
}
