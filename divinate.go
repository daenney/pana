package pana

import (
	"cmp"
	"encoding/json"
	"maps"
	"slices"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/vocab"
	"sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Divinate attempts to construct the minimal compaction context to compact all
// known properties in a document.
//
// This is done using the knowledge of the vocab packages. As such, it cannot
// determine a compaction term definition for a term Pana doesn't recognise.
func Divinate(doc ld.Node) (json.RawMessage, error) {
	remotes := map[string]struct{}{
		activitystreams.IRI: {},
	}

	local := map[string]any{}

	divinate(doc, remotes, local)

	lr := len(remotes)
	le := len(local)

	if lr == 0 && le == 0 {
		return nil, nil
	}

	if lr == 1 && le == 0 {
		return json.RawMessage(`"` + slices.Collect(maps.Keys(remotes))[0] + `"`), nil
	}

	rems := slices.Collect(maps.Keys(remotes))
	slices.SortFunc(rems, func(a string, b string) int {
		if a == activitystreams.IRI {
			return -1
		}
		if b == activitystreams.IRI {
			return 1
		}
		return cmp.Compare(a, b)
	})

	data := make([]any, 0, lr+1)
	for _, r := range rems {
		data = append(data, r)
	}

	if le != 0 {
		data = append(data, local)
	}

	return json.Marshal(data)
}

func divinate(doc ld.Node, remotes map[string]struct{}, elems map[string]any) {
	if doc.ID != "" {
		lookup(doc.ID, remotes, elems)
	}

	for prop, val := range doc.Properties {
		lookup(prop, remotes, elems)

		for _, n := range val {
			divinate(n, remotes, elems)
		}
	}

	for _, n := range doc.List {
		divinate(n, remotes, elems)
	}

	for _, t := range doc.Type {
		lookup(t, remotes, elems)
	}
}

func lookup(iri string, remotes map[string]struct{}, elems map[string]any) {
	def := vocab.Lookup(iri)
	if def == nil {
		return
	}

	switch data := def.(type) {
	case string:
		remotes[data] = struct{}{}
	case map[string]any:
		maps.Copy(elems, data)
	default:
		panic("vocab.Lookup returned nonesense")
	}
}
