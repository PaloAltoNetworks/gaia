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
	defer file.Close()

	var decodedTypeMapping map[string]struct {
		OpenAPI3 struct {
			Type interface{} `yaml:"type,omitempty"`
		} `yaml:"openapi3,omitempty"`
	}
	if err := yaml.NewDecoder(file).Decode(&decodedTypeMapping); err != nil {
		t.Fatalf("error decoding the content of _type.mapping file: %v", err)
	}

	for typeName, mapping := range decodedTypeMapping {
		if mapping.OpenAPI3.Type == nil {
			t.Errorf("'%s' has no type mapping defined for 'openapi3'", typeName)
		}
	}
}
