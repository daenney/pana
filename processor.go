package pana

import (
	"fmt"
	"log/slog"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/internal/json"
	"sourcery.dny.nu/pana/internal/loader"
	"sourcery.dny.nu/pana/vocab/schema"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
)

// Processor is used to process incoming messages and prepare outgoing messages.
//
// Your application only needs a single processor instance.
type Processor struct {
	ldproc *ld.Processor
	l      *loader.Builtin
}

// NewProcessor initialised a new [Processor] suitable for dealing with
// ActivityStreams messages.
//
// It uses [loader.Builtin] to retrieve contexts and does not retrieve them over
// the network.
func NewProcessor(
	logger *slog.Logger,
) *Processor {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	loader := loader.New()
	return &Processor{
		l: loader,
		ldproc: ld.NewProcessor(
			ld.WithLogger(logger),
			ld.WithCompactArrays(true),
			ld.WithCompactToRelative(false), // Avoids compacting to relative IRIs
			ld.With10Processing(false),      // Misskey seems to do some JSON-LD 1.1
			ld.WithRemoteContextLoader(loader.Get),
			ld.WithExcludeIRIsFromCompaction(as.PublicCollection),
			ld.WithRemapPrefixIRIs("http://schema.org#", schema.Namespace),
		),
	}
}

// Marshal takes an [Activity] and returns JSON-LD in compacted document form.
//
// This is the shape of JSON you want to exchange with other servers and
// clients.
//
// The compaction context is a JSON document representing the JSON-LD context
// that should be used for compaction.
func (p *Processor) Marshal(
	compactionContext json.RawMessage,
	activity Activity,
) (json.RawMessage, error) {
	if compactionContext[0] == '{' {
		ctx, err := json.GetContextDocument(compactionContext)
		if err != nil {
			return nil, err
		}
		compactionContext = ctx
	}

	return p.ldproc.Compact(
		compactionContext,
		[]ld.Node{ld.Node(activity)},
		"",
	)
}

// Unmarshal takes a JSON document and returns an [Activity] that represents it
// using JSON-LD expanded document form.
//
// In JSON-LD the result of expanding a document is always a list. But in the
// case of ActivityPub we only ever send out a single document at a time, so it
// returns [Activity]. If the result happens to have more than one object an
// error is raised so it doesn't go unnoticed.
func (p *Processor) Unmarshal(
	document json.RawMessage,
) (*Activity, error) {
	res, err := p.ldproc.Expand(document, "")
	if err != nil {
		return nil, err
	}

	if lres := len(res); lres != 1 {
		return nil, fmt.Errorf("expected a single document, got: %d", lres)
	}

	return (*Activity)(&res[0]), nil
}

// RegisterContextURL adds or overrides a context document for the specified
// remote context URL in the loader.
func (p *Processor) RegisterContextURL(
	url string,
	doc json.RawMessage,
) error {
	return p.l.RegisterContextURL(url, doc)
}

// RegisterContextPath adds or overrides a context document for the specified
// remote path in the loader.
//
// Paths are always matches as URL suffixes.
func (p *Processor) RegisterContextPath(
	path string,
	doc json.RawMessage,
) error {
	return p.l.RegisterContextPath(path, doc)
}
