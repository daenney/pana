package pana

import (
	ld "sourcery.dny.nu/longdistance"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Article is the ActivityStreams Article type.
type Article Object

// NewArticle initialises a new Article.
func NewArticle() *Article {
	return &Article{
		Properties: make(ld.Properties),
		Type:       []string{as.TypeArticle},
	}
}

// Build finalises the Article.
//
// This returns [Any] since that's what [Activity.SetObject] expects.
func (a *Article) Build() Any {
	return Any(*a)
}
