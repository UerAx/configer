package configer

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var (
	commonError error
)

func init() {
	commonError = errors.New("section or key no exist")
}

type ConfigFile struct {
	data map[string]map[string]string // Section -> key : value
	path string
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
		data: make(map[string]map[string]string, 0),
		path: fileName,
	}
}

func (cfg *ConfigFile) load(fileName string) error {
	cfg.path = fileName
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
			cfg.data[key] = make(map[string]string, 10)
		}
		if len(key) < 1 {
			continue
		}
		for i, v := range line {
			if v == '=' {
				cfg.setValue(key, string(line[:i]), string(line[i+1:]))
				break
			}
		}
	}
	return nil
}

func (cfg *ConfigFile) setValue(section, key, value string) {
	cfg.data[section][key] = value
}

func (cfg *ConfigFile) getValue(section, key string) (string, error) {
	if _, ok := cfg.data[section][key]; !ok {
		return "", commonError
	}
	return cfg.data[section][key], nil
}
