package pana

import (
	ld "sourcery.dny.nu/longdistance"
)

// Has checks if an object has a specific property set.
//
// It handles JSON-LD keyword aliasses for id and type.
func Has[T node](in *T, property string) bool {
	if in == nil {
		return false
	}

	if property == "id" {
		property = ld.KeywordID
	}

	if property == "type" {
		property = ld.KeywordType
	}

	n := ld.Node(*in)
	return n.Has(property)
}

// IsReference indicates if this object is a reference.
//
// This means it only has the ID, and optionally a Type, set. You'll need to
// retrieve the object using the ID to get additional properties.
func IsReference[T node](in *T) bool {
	if in == nil {
		return false
	}

	n := ld.Node(*in)
	return n.IsSubjectReference()
}

// IsObject indicates if this object is a (partially) complete object.
//
// This means it has an ID, optionally a Type and at least one other
// property. It doesn't mean the object representation is complete, and you may
// need to retrieve the object using the ID to get additional properties.
func IsObject[T node](in *T) bool {
	if in == nil {
		return false
	}

	n := ld.Node(*in)
	return n.IsSubject()
}

// Properties returns a set with an entry for each property set on an object.
//
// It handles JSON-LD keyword aliasses for id and type.
func Properties[T node](in *T) map[string]struct{} {
	if in == nil {
		return nil
	}

	n := ld.Node(*in)
	s := n.PropertySet()

	if _, ok := s[ld.KeywordID]; ok {
		delete(s, ld.KeywordID)
		s["id"] = struct{}{}
	}

	if _, ok := s[ld.KeywordType]; ok {
		delete(s, ld.KeywordType)
		s["type"] = struct{}{}
	}

	return s
}
