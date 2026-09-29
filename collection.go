package pana

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"iter"
	"slices"

	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	"sourcery.dny.nu/pana/vocab/w3/xmlschema"
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

// NewOrderedCollection initialises a new Collection with
// [as.TypeOrderedCollection].
func NewOrderedCollection() *Collection {
	return &Collection{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeOrderedCollection},
	}
}

// Build finalises the Collection.
func (c *Collection) Build() Collection {
	return *c
}

// IsOrdered checks if this is an ordered collection.
func (c *Collection) IsOrdered() bool {
	if c == nil {
		return false
	}

	return slices.Contains(c.Type, as.TypeOrderedCollection)
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

// GetItems returns the values in [as.Items].
//
// This returns Any because a collection can contain objects of any type.
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-items.
func (c *Collection) GetItems() iter.Seq[*Any] {
	return getCollectionItem((*ld.Node)(c), c.IsOrdered())
}

// AddItems appends items to [as.Items].
func (c *Collection) AddItems[T node](items ...T) *Collection {
	addCollectionItem((*ld.Node)(c), c.IsOrdered(), items)
	return c
}

// GetTotalItems returns the value from [as.TotalItems]
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-totalitems.
func (c *Collection) GetTotalItems() jsontext.Value {
	if nodes := (*ld.Node)(c).GetNodes(as.TotalItems); len(nodes) == 1 {
		return nodes[0].Value
	}

	return nil
}

// SetTotalItems sets the value in [as.TotalItems].
func (c *Collection) SetTotalItems(v uint) *Collection {
	data, _ := json.Marshal(v)
	(*ld.Node)(c).SetNodes(as.TotalItems, ld.Node{Value: data, Type: []string{xmlschema.TypeNonNegativeInteger}})
	return c
}

// SetTotalItemsRaw sets the value in [as.TotalItems].
func (c *Collection) SetTotalItemsRaw(v jsontext.Value) *Collection {
	(*ld.Node)(c).SetNodes(as.TotalItems, ld.Node{Value: v, Type: []string{xmlschema.TypeNonNegativeInteger}})
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

// NewOrderedCollectionPage initialises a new CollectionPage with
// [as.TypeOrderedCollectionPage].
func NewOrderedCollectionPage() *CollectionPage {
	return &CollectionPage{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeOrderedCollectionPage},
	}
}

// Build finalises the CollectionPage.
func (p *CollectionPage) Build() CollectionPage {
	return *p
}

// IsOrdered checks if this is an ordered collection page.
func (p *CollectionPage) IsOrdered() bool {
	if p == nil {
		return false
	}

	return slices.Contains(p.Type, as.TypeOrderedCollectionPage)
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

// GetItems returns the values in [as.Items].
//
// This returns Any because a collection page can contain objects of any type.
//
// See https://www.w3.org/TR/activitystreams-vocabulary/#dfn-items.
func (p *CollectionPage) GetItems() iter.Seq[*Any] {
	return getCollectionItem((*ld.Node)(p), p.IsOrdered())
}

// AddItems appends items to [as.Items].
func (p *CollectionPage) AddItems[T node](items ...T) *CollectionPage {
	addCollectionItem((*ld.Node)(p), p.IsOrdered(), items)
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
