package jsonutil

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
)

func GetContextDocument(doc jsontext.Value) (jsontext.Value, error) {
	var check struct {
		Context jsontext.Value `json:"@context,omitzero"`
	}

	if err := json.Unmarshal(doc, &check); err != nil {
		return nil, err
	}

	if check.Context != nil {
		return check.Context, nil
	}

	return doc, nil
}
