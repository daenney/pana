package pana

import ld "sourcery.dny.nu/longdistance"

// Any is returned when a property can return [Link] or [Object].
//
// You can use [Any.GetType] to determine what to cast it to.
type Any ld.Node

// See [Object.GetType].
func (a *Any) GetType() string {
	return (*Object)(a).GetType()
}
