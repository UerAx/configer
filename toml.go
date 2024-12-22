package goconf

import "github.com/BurntSushi/toml"

type TOMLParser struct {}

func (t *TOMLParser) Parse(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}

func (t *TOMLParser) FileExt() []string {
	return []string{"toml"}
}