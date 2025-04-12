package json

import "encoding/json"

type RawMessage = json.RawMessage

func GetContextDocument(doc RawMessage) (RawMessage, error) {
	var check struct {
		Context RawMessage `json:"@context,omitempty"`
	}

	if err := json.Unmarshal(doc, &check); err != nil {
		return nil, err
	}

	if check.Context != nil {
		return check.Context, nil
	}

	return doc, nil
}
