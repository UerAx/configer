package goconf

import (
	"errors"
)

var cfg *CfgFile

func init() {
	cfg = &CfgFile{Data: make(map[string]interface{}, 0), tmp: make(map[string]interface{}, 0)}
}

// 直接获取CfgFile对象, 用于需要自己构建结构体的情况下
// 将map[filename]的interface{}转成自己需要的结构体
func NewConfigFile() *CfgFile {
	return &CfgFile{Data: make(map[string]interface{}, 0), tmp: make(map[string]interface{}, 0)}
}

func LoadConfig(path string) error {
	return cfg.ReadFiles(path)
}

func VarStringOrDefault(def string, keys ...string) string {
	val, err := VarString(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarString(keys ...string) (string, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return "", errors.New("未找到对应的值")
	}
	if v, ok := val.(string); !ok {
		return "", errors.New("无法转换成string")
	} else {
		return v, nil
	}
}

func VarInt64(keys ...string) (int64, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(int64); !ok {
		return 0, errors.New("无法转换成int64")
	} else {
		return v, nil
	}
}

func VarInt64OrDefault(def int64, keys ...string) int64 {
	val, err := VarInt64(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarInt32(keys ...string) (int32, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(int32); !ok {
		return 0, errors.New("无法转换成int32")
	} else {
		return v, nil
	}
}

func VarInt32OrDefault(def int32, keys ...string) int32 {
	val, err := VarInt32(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarUint32(keys ...string) (uint32, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(uint32); !ok {
		return 0, errors.New("无法转换成uint32")
	} else {
		return v, nil
	}
}

func VarUint32OrDefault(def uint32, keys ...string) uint32 {
	val, err := VarUint32(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarUint64(keys ...string) (uint64, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(uint64); !ok {
		return 0, errors.New("无法转换成uint64")
	} else {
		return v, nil
	}
}

func VarUint64OrDefault(def uint64, keys ...string) uint64 {
	val, err := VarUint64(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarBool(keys ...string) (bool, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return false, errors.New("未找到对应的值")
	}
	if v, ok := val.(bool); !ok {
		return false, errors.New("无法转换成bool")
	} else {
		return v, nil
	}
}

func VarBoolOrDefault(def bool, keys ...string) bool {
	val, err := VarBool(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarArray(keys ...string) ([]interface{}, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return nil, errors.New("未找到对应的值")
	}
	if v, ok := val.([]interface{}); !ok {
		return nil, errors.New("无法转换成数组")
	} else {
		return v, nil
	}
}

func VarArrayInt(keys ...string) ([]int, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]int, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(int); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayInt32(keys ...string) ([]int32, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]int32, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(int32); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayInt64(keys ...string) ([]int64, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]int64, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(int64); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayString(keys ...string) ([]string, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]string, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(string); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayFloat32(keys ...string) ([]float32, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]float32, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(float32); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayFloat64(keys ...string) ([]float64, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]float64, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(float64); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayUInt32(keys ...string) ([]uint32, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]uint32, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(uint32); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayUInt64(keys ...string) ([]uint64, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]uint64, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(uint64); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarArrayUInt(keys ...string) ([]uint, error) {
	val, err := VarArray(keys...)
	if err != nil {
		return nil, err
	}

	conv := make([]uint, 0, len(val))
	
	for _, v := range val {
		if i, ok := v.(uint); ok {
			conv = append(conv, i)
		}
	}
	return conv, nil
}

func VarMap(keys ...string) (map[string]interface{}, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return nil, errors.New("未找到对应的值")
	}
	if v, ok := val.(map[string]interface{}); !ok {
		return nil, errors.New("无法转换成Map")
	} else {
		return v, nil
	}
}

func VarInt(keys ...string) (int, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(int); !ok {
		return 0, errors.New("无法转换成int")
	} else {
		return v, nil
	}
}

func VarIntOrDefault(def int, keys ...string) int {
	val, err := VarInt(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarFloat32(keys ...string) (float32, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(float32); !ok {
		return 0, errors.New("无法转换成float32")
	} else {
		return v, nil
	}
}

func VarFloat32OrDefault(def float32, keys ...string) float32 {
	val, err := VarFloat32(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarFloat64(keys ...string) (float64, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(float64); !ok {
		return 0, errors.New("无法转换成float64")
	} else {
		return v, nil
	}
}

func VarFloat64OrDefault(def float64, keys ...string) float64 {
	val, err := VarFloat64(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarUint(keys ...string) (uint, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(uint); !ok {
		return 0, errors.New("无法转换成uint")
	} else {
		return v, nil
	}
}

func VarUintOrDefault(def uint, keys ...string) uint {
	val, err := VarUint(keys...)
	if err != nil {
		return def
	}
	return val
}

func VarRune(keys ...string) (rune, error) {
	val := cfg.GetValue(keys...)
	if val == nil {
		return 0, errors.New("未找到对应的值")
	}
	if v, ok := val.(rune); !ok {
		return 0, errors.New("无法转换成rune")
	} else {
		return v, nil
	}
}

func VarRuneOrDefault(def rune, keys ...string) rune {
	val, err := VarRune(keys...)
	if err != nil {
		return def
	}
	return val
}