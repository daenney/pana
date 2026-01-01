package pana

import (
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"

	ld "sourcery.dny.nu/longdistance"
	"sourcery.dny.nu/pana/internal/json"
	"sourcery.dny.nu/pana/internal/loader"
	as "sourcery.dny.nu/pana/vocab/w3/activitystreams"
	secv1 "sourcery.dny.nu/pana/vocab/w3id/securityv1"
)

// Processor is used to process incoming messages and prepare outgoing messages.
//
// Your application only needs a single processor instance.
type Processor struct {
	ldproc *ld.Processor
}

// NewProcessor initialised a new [Processor] suitable for dealing with
// ActivityStreams messages.
//
// It uses [loader.Builtin] to retrieve contexts and does not retrieve them over
// the network. It also uses [ValidateContext] to enforce additional constraints
// on the context term definitions.
func NewProcessor(
	logger *slog.Logger,
) *Processor {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}

	return &Processor{
		ldproc: ld.NewProcessor(
			ld.WithLogger(logger),
			ld.WithCompactArrays(true),
			ld.WithCompactToRelative(false), // Avoids compacting to relative IRIs
			ld.With10Processing(false),      // Misskey seems to do some JSON-LD 1.1
			ld.WithRemoteContextLoader((loader.New()).Get),
			ld.WithExcludeIRIsFromCompaction(as.PublicCollection),
			ld.WithValidateContext(ValidateContext),
		),
	}
}

// Marshal takes an [Any] and returns JSON-LD in compacted document form.
//
// This is the shape of JSON you want to exchange with other servers and
// clients.
//
// The compaction context is a JSON document representing the JSON-LD context
// that should be used for compaction. If the compaction context is not
// provided, [Divinate] will be used to attempt and determine an appropriate
// one.
func (p *Processor) Marshal(
	dst io.Writer,
	compactionContext json.RawMessage,
	object Any,
) error {
	if compactionContext == nil {
		res, err := Divinate(ld.Node(object))
		if err != nil {
			return fmt.Errorf("context divination failed: %w", err)
		}
		compactionContext = res
	}

	if compactionContext[0] == '{' {
		ctx, err := json.GetContextDocument(compactionContext)
		if err != nil {
			return err
		}
		compactionContext = ctx
	}

	return p.ldproc.Compact(
		dst,
		compactionContext,
		[]ld.Node{ld.Node(object)},
		"",
	)
}

// Unmarshal takes a JSON document and returns an [Activity] that represents it
// using JSON-LD expanded document form.
//
// If the document was retrieved over HTTP, the request URL should be passed
// as the second argument.
//
// In JSON-LD the result of expanding a document is always a list. But in the
// case of ActivityPub we only ever send out a single document at a time, so it
// returns [Any]. If the result happens to have more than one object an error is
// raised so it doesn't go unnoticed.
func (p *Processor) Unmarshal(
	document io.Reader,
	url string,
) (*Any, error) {
	res, err := p.ldproc.Expand(document, url)
	if err != nil {
		return nil, err
	}

	if lres := len(res); lres != 1 {
		return nil, fmt.Errorf("expected a single document, got: %d", lres)
	}

	return (*Any)(&res[0]), nil
}

// ValidateContext enforces that the term to IRI mapping for certain contexts
// match their normative context definition. This ensures IRIs aren't compacted
// to unexpected terms.
func ValidateContext(ctx *ld.Context) bool {
	for term, def := range ctx.Terms() {
		if strings.HasPrefix(def.IRI, as.Namespace) {
			if def.Prefix {
				if term == "as" && def.IRI != as.Namespace {
					return false
				}
				continue
			}

			short := as.Term(def.IRI)

			if slices.Equal(def.Container, []string{ld.KeywordLanguage}) {
				if term != short && term != short+"Map" {
					return false
				}
				continue
			}

			if slices.Equal(def.Container, []string{ld.KeywordList}) {
				if term != short && term != "ordered"+strings.ToUpper(short[:1])+short[1:] {
					return false
				}
				continue
			}

			if term != as.Term(def.IRI) {
				return false
			}

			continue
		}

		if strings.HasPrefix(def.IRI, secv1.Namespace) {
			if def.Prefix {
				if term == "sec" && def.IRI != secv1.Namespace {
					return false
				}
				continue
			}

			short := secv1.Term(def.IRI)
			if term != short {
				return false
			}
		}
	}

	return true
}
