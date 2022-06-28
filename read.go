package configer

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var (
	commonError error
)

func init() {
	commonError = errors.New("section or key no exist")
}

type ConfigFile struct {
	Data map[string]map[string]string // Section -> key : value
	Path string
}

func LoadConfigFile(fileName string) (*ConfigFile, error) {
	cfg := newConfig(fileName)
	err := cfg.load(fileName)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func newConfig(fileName string) *ConfigFile {
	return &ConfigFile{
		Data: make(map[string]map[string]string, 0),
		Path: fileName,
	}
}

func (cfg *ConfigFile) load(fileName string) error {
	cfg.Path = fileName
	data := make(map[string]map[string]string, 0)
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer func() {
		file.Close()
	}()
	reader := bufio.NewReader(file)
	var key string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if len(line) == 0 {
			continue
		}
		if line[0] == '[' && len(line) > 1 && line[len(line)-1] == ']' {
			key = string(line[1 : len(line)-1])
			data[key] = make(map[string]string, 10)
		}
		idx := strings.Index(string(line), "=")
		if idx > 0 {
			v := string(line[idx+1:])
			if v[0] == '"' && v[len(v) - 1] == '"' {
				v = v[1:len(v)-1]
			}
			data[key][string(line[:idx])] = v
		}
	}
	cfg.Data = data
	return nil
}

func (cfg *ConfigFile) getValue(section, key string) (string, error) {
	if _, ok := cfg.Data[section][key]; !ok {
		return "", commonError
	}
	return cfg.Data[section][key], nil
}
