package gaia

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestTypeMapping_all_have_key_openapi3(t *testing.T) {

	file, err := os.Open("./specs/_type.mapping")
	if err != nil {
		t.Fatalf("error opening file for reading: %v", err)
	}

	var decodedTypeMapping map[string]struct {
		OpenAPI3 interface{} `yaml:"openapi3,omitempty"`
	}
	if err := yaml.NewDecoder(file).Decode(&decodedTypeMapping); err != nil {
		t.Fatalf("error decoding the content of _type.mapping file: %v", err)
	}

	for typeName, mapping := range decodedTypeMapping {
		if mapping.OpenAPI3 == nil {
			t.Errorf("type '%s' has no mapping defined for 'openapi3'", typeName)
		}
	}
}
