package pana

import (
	"unsafe"

	ld "code.dny.dev/longdistance"
)

// Has checks if an object has a specific property set.
//
// It handles JSON-LD keyword aliasses for id and type.
func Has[T ld.Internal](in *T, property string) bool {
	if property == "id" {
		property = ld.KeywordID
	}

	if property == "type" {
		property = ld.KeywordType
	}

	return (*ld.Node)(unsafe.Pointer(in)).Has(property)
}

// IsReference indicates if this object is a reference.
//
// This means it only has the ID, and optionally a Type, set. You'll need to
// retrieve the object using the ID to get additional properties.
func IsReference[T ld.Internal](in *T) bool {
	return (*ld.Node)(unsafe.Pointer(in)).IsSubjectReference()
}

// IsObject indicates if this object is a (partially) complete object.
//
// This means it has an ID, optionally a Type and at least one other
// property. It doesn't mean the object representation is complete, and you may
// need to retrieve the object using the ID to get additional properties.
func IsObjec[T ld.Internal](in *T) bool {
	return (*ld.Node)(unsafe.Pointer(in)).IsSubject()
}

// Properties returns a set with an entry for each property set on an object.
//
// It handles JSON-LD keyword aliasses for id and type.
func Properties[T ld.Internal](in *T) map[string]struct{} {
	s := (*ld.Node)(unsafe.Pointer(in)).PropertySet()

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
