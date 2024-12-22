package goconf

import (
	"fmt"
	"os"
	"strings"
)

type ConfigParser interface {
    Parse(data []byte, v any) error
    FileExt() []string // 返回支持的文件扩展名
}

type ConfigManager struct {
	parsers map[string]ConfigParser
    path string
}

func newConfigManager() *ConfigManager {
    return &ConfigManager{
        parsers: make(map[string]ConfigParser),
        path: "",
    }
}

// 自动注册默认解析器
func DefaultConfigManager() *ConfigManager {
    manager := newConfigManager()
    manager.RegisterParser(&JSONParser{})
    manager.RegisterParser(&TOMLParser{})
    manager.RegisterParser(&YAMLParser{})
    return manager
}

func (t *ConfigManager) RegisterParser(parser ConfigParser) {
	for _, ext := range parser.FileExt() {
		t.parsers[ext] = parser
	}
}

func (cm *ConfigManager) Load(filePath string, v interface{}) error {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return fmt.Errorf("file does not exist: %s", filePath)
    }

	data, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }

    cm.path = filePath

    ext := ""
	if idx := strings.LastIndex(filePath, "."); idx != -1 {
		ext = filePath[idx+1:]
	}

    parser, exists := cm.parsers[ext]
    if !exists {
        return fmt.Errorf("unsupported file type: %s", ext)
    }

    return parser.Parse(data, v)
}

func (cm *ConfigManager) Reload(filePath string, v interface{}) error {
    return cm.Load(filePath, v)
}