// internal/config/config.go

package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Field represents a field within an entity.
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// Entity represents an entity with a name and a set of fields.
type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

// Manifest represents the root structure of the YAML configuration file.
type Manifest struct {
	Entities []Entity `yaml:"entities"`
}

// LoadManifest reads and unmarshals the specified YAML file into a Manifest struct.
func LoadManifest(filename string) (*Manifest, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}
