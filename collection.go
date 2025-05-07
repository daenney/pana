package pana

import (
	"encoding/json"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Collection is the ActivityStreams Collection type.
//
// It is also used for the OrderedCollection. Use [Collection.IsOrdered] to see,
// or [Collection.GetType].
type Collection Object

// NewCollection initialises a new Collection.
func NewCollection() *Collection {
	return &Collection{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeCollection},
	}
}

// Build finalises the Collection.
func (c *Collection) Build() Collection {
	return *c
}

// See [Object.GetID].
func (c *Collection) GetID() string {
	return (*Object)(c).GetID()
}

// See [Object.SetID].
func (c *Collection) SetID(id string) *Collection {
	(*Object)(c).SetID(id)
	return c
}

// See [Object.GetType].
func (c *Collection) GetType() string {
	return (*Object)(c).GetType()
}

// See [Object.SetType].
func (c *Collection) SetType(typ string) *Collection {
	(*Object)(c).SetType(typ)
	return c
}

// GetFirst returns the [CollectionPage] in [as.First].
func (c *Collection) GetFirst() *CollectionPage {
	if nodes := (*ld.Node)(c).GetNodes(as.First); len(nodes) == 1 {
		return (*CollectionPage)(&nodes[0])
	}

	return nil
}

// SetFirst sets the [CollectionPage] in [as.First].
func (c *Collection) SetFirst(p CollectionPage) *Collection {
	(*ld.Node)(c).SetNodes(as.First, ld.Node(p))
	return c
}

// GetTotalItems returns the value from [as.TotalItems]
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-totalitems.
func (c *Collection) GetTotalItems() json.RawMessage {
	if nodes := (*ld.Node)(c).GetNodes(as.TotalItems); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetTotalItems sets the value in [as.TotalItems].
func (c *Collection) SetTotalItems(v json.RawMessage) *Collection {
	(*ld.Node)(c).SetNodes(as.TotalItems, ld.Node{Value: v})
	return c
}

// CollectionPage is the ActivityStreams CollectionPage type.
//
// It is also used for the OrderedCollectionPage. Use [CollectionPage.IsOrdered]
// to see, or [CollectionPage.GetType].
type CollectionPage Object

// NewCollectionPage initialises a new CollectionPage.
func NewCollectionPage() *CollectionPage {
	return &CollectionPage{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeCollectionPage},
	}
}

// Build finalises the CollectionPage.
func (p *CollectionPage) Build() CollectionPage {
	return *p
}

// See [Object.GetType].
func (p *CollectionPage) GetType() string {
	return (*Object)(p).GetType()
}

// See [Object.SetType].
func (p *CollectionPage) SetType(typ string) *CollectionPage {
	(*Object)(p).SetType(typ)
	return p
}

// GetNext returns the URL stored in [as.Next].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-next.
func (p *CollectionPage) GetNext() string {
	if nodes := (*ld.Node)(p).GetNodes(as.Next); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetNext sets the URL in [as.Next].
func (p *CollectionPage) SetNext(url string) *CollectionPage {
	(*ld.Node)(p).SetNodes(as.Next, ld.Node{ID: url})
	return p
}

// GetPartOf returns the URL stored in [as.PartOf].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-partof.
func (p *CollectionPage) GetPartOf() string {
	if nodes := (*ld.Node)(p).GetNodes(as.PartOf); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetPartOf sets the URL in [as.PartOf].
func (p *CollectionPage) SetPartOf(url string) *CollectionPage {
	(*ld.Node)(p).SetNodes(as.PartOf, ld.Node{ID: url})
	return p
}

// GetPrev returns the URL stored in [as.Prev].
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-prev.
func (p *CollectionPage) GetPrev() string {
	if nodes := (*ld.Node)(p).GetNodes(as.Prev); len(nodes) == 1 {
		return nodes[0].ID
	}

	return ""
}

// SetPrev sets the URL in [as.Prev].
func (p *CollectionPage) SetPrev(url string) *CollectionPage {
	(*ld.Node)(p).SetNodes(as.Prev, ld.Node{ID: url})
	return p
}
