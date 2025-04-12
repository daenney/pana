package pana

import ld "code.dny.dev/longdistance"

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

func toLDNodes[T ld.Internal](ins ...T) []ld.Node {
	if len(ins) == 0 {
		return nil
	}

	res := make([]ld.Node, 0, len(ins))
	for _, in := range ins {
		res = append(res, ld.Node(in))
	}

	return res
}
