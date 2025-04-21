package loader_test

import (
	"bytes"
	"errors"
	"testing"

	ld "code.dny.dev/longdistance"
	"code.dny.dev/pana/internal/json"
	"code.dny.dev/pana/internal/loader"
	"code.dny.dev/pana/vocab/litepub"
	"code.dny.dev/pana/vocab/w3/activitystreams"
)

func TestLoaderURL(t *testing.T) {
	l := loader.New()

	data, err := l.Get(t.Context(), activitystreams.IRI)
	if err != nil {
		t.Fatalf("expected no error, got: %s", err)
	}

	ctxdoc, err := json.GetContextDocument(activitystreams.ContextDocument)
	if err != nil {
		t.Fatalf("expected document to have @context, got: %s", err)
	}

	if !bytes.Equal(data.Context, ctxdoc) {
		t.Fatal("contexts should be equal")
	}
}

func TestLoaderPath(t *testing.T) {
	l := loader.New()
	data, err := l.Get(t.Context(), "http://example.org/some/path/litepub-0.1.jsonld")
	if err != nil {
		t.Fatalf("expected no error, got: %s", err)
	}

	ctxdoc, err := json.GetContextDocument(litepub.ContextDocument1dot0)
	if err != nil {
		t.Fatalf("expected document to have @context, got: %s", err)
	}

	if !bytes.Equal(data.Context, ctxdoc) {
		t.Fatal("contexts should be equal")
	}
}

func TestLoaderError(t *testing.T) {
	l := loader.New()
	_, err := l.Get(t.Context(), "http://example.com/unknown")

	if !errors.Is(err, ld.ErrLoadingRemoteContext) {
		t.Fatalf("expected loading remote context error, got: %s", err)
	}
}
