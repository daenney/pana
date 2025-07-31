package pana

import (
	"encoding/json"

	ld "sourcery.dny.nu/longdistance"
)

// Localised represents a value with an optional language tag and direction.
//
// It mirrors a language-mapped attribute in JSON-LD.
type Localised ld.Node
type Localized = Localised

// NewLocalised initialises a new Localised.
func NewLocalised() *Localised {
	return &Localised{}
}

// Build finalises the Localised.
func (l *Localised) Build() Localised {
	return *l
}

// GetLanguage returns a normalised to lower-case BCP-47 language tag.
//
// The empty string indicates that the language is unknown, not English.
func (l *Localised) GetLanguage() string {
	return l.Language
}

// SetLanguage sets the language tag for the value.
//
// This must be a valid BCP-47 language tag and may be normalised to lower-case.
func (l *Localised) SetLanguage(lang string) *Localised {
	l.Language = lang
	return l
}

// GetValue returns the value.
func (l *Localised) GetValue() json.RawMessage {
	return l.Value
}

// SetValue sets the strings as the value.
//
// The string must be valid UTF-8. This will always be the case unless the
// string was directly constructed from a []byte. Use
// [unicode/utf8.ValidString] to check for that if necessary.
func (l *Localised) SetValue(value string) *Localised {
	data, _ := json.Marshal(value)
	l.Value = data
	return l
}

// SetValueRaw sets the value.
//
// The value must be a valid UTF-8 byte sequence. Use [unicode/utf8.Valid] to
// check for that if necessary.
func (l *Localised) SetValueRaw(value json.RawMessage) *Localised {
	l.Value = value
	return l
}
