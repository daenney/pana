package loader

import (
	"context"
	"maps"
	"slices"
	"strings"
	"sync"

	ld "code.dny.dev/longdistance"
	"code.dny.dev/pana/internal/json"
	"code.dny.dev/pana/vocab/geojson"
	gts "code.dny.dev/pana/vocab/gotosocial"
	"code.dny.dev/pana/vocab/litepub"
	as "code.dny.dev/pana/vocab/w3/activitystreams"
	credv1 "code.dny.dev/pana/vocab/w3id/credentialsv1"
	idv1 "code.dny.dev/pana/vocab/w3id/identityv1"
	secv1 "code.dny.dev/pana/vocab/w3id/securityv1"
)

// Builtin is a remote context cache.
type Builtin struct {
	domains  map[string]ld.Document
	paths    map[string]ld.Document
	pathKeys []string

	lock sync.RWMutex
}

func New() *Builtin {
	b := &Builtin{
		domains: make(map[string]ld.Document, 8),
		paths:   make(map[string]ld.Document, 2),
	}
	b.RegisterContextURL(as.IRI, as.ContextDocument)
	b.RegisterContextURL(gts.IRI, gts.ContextDocument)
	b.RegisterContextURL(credv1.IRI, credv1.ContextDocument)
	b.RegisterContextURL(idv1.IRI, idv1.ContextDocument)
	b.RegisterContextURL(secv1.IRI, secv1.ContextDocument)
	b.RegisterContextURL(litepub.IRI, litepub.ContextDocument)
	b.RegisterContextURL(geojson.IRI, geojson.ContextDocument)
	b.RegisterContextPath("/litepub-0.1.jsonld", litepub.ContextDocument1dot0)
	return b
}

// Get is an [ld.RemoteContextLoaderFunc].
func (b *Builtin) Get(_ context.Context, url string) (ld.Document, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	if doc, ok := b.domains[url]; ok {
		return doc, nil
	}

	for _, path := range b.pathKeys {
		if strings.HasSuffix(url, path) {
			doc := b.paths[path]
			doc.URL = url
			return doc, nil
		}
	}

	return ld.Document{}, ld.ErrLoadingRemoteContext
}

// RegisterContextURL adds or overrides a context document for the specified
// remote context URL in the loader.
func (b *Builtin) RegisterContextURL(url string, doc json.RawMessage) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	ctx, err := json.GetContextDocument(doc)
	if err != nil {
		return err
	}

	b.domains[url] = ld.Document{URL: url, Context: ctx}
	return nil
}

// RegisterContextPath adds or overrides a context document for the specified
// remote path in the loader.
//
// Paths are always matches as URL suffixes.
func (b *Builtin) RegisterContextPath(path string, doc json.RawMessage) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	ctx, err := json.GetContextDocument(doc)
	if err != nil {
		return err
	}

	b.paths[path] = ld.Document{URL: "", Context: ctx}
	b.pathKeys = slices.Collect(maps.Keys(b.paths))
	return nil
}
