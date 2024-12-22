package goconf

import "encoding/json"


type JSONParser struct {}

func (t *JSONParser) Parse(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (t *JSONParser) FileExt() []string {
	return []string{"json"}
}