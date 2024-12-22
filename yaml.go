package goconf

import (
	"gopkg.in/yaml.v2"
)

type YAMLParser struct {}

func (t *YAMLParser) Parse(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (t *YAMLParser) FileExt() []string {
	return []string{"yaml", "yml"}
}