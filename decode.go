package initapp

import (
	"encoding/json"

	yaml "github.com/goccy/go-yaml"
)

// decodeJSON converts JSON to a structure
func decodeJSON[T interface{}](raw []byte) (*T, error) {
	var data T
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func decodeYAML[T interface{}](raw []byte) (*T, error) {
	var data T
	if err := yaml.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
