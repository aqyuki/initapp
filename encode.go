package initapp

import (
	"encoding/json"

	yaml "github.com/goccy/go-yaml"
)

// encodeJSON convert structure to a JSON
func encodeJSON[T interface{}](normal *T) ([]byte, error) {
	var b []byte
	b, err := json.Marshal(normal)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// encodeYAML convert structure to a YAML
func encodeYAML[T interface{}](normal *T) ([]byte, error) {
	var b []byte
	b, err := yaml.Marshal(normal)
	if err != nil {
		return nil, err
	}
	return b, nil
}
