package configer

import "strconv"

func (cfg *ConfigFile) GetStringValueOrDefault(section, key, defaultVal string) string {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}
	return value
}

func (cfg *ConfigFile) GetIntValueOrDefault(section, key string, defaultVal int) int {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}
	res, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}
	return res
}

func (cfg *ConfigFile) GetInt64ValueOrDefault(section, key string, defaultVal int64) int64 {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultVal
	}
	return res
}

func (cfg *ConfigFile) GetInt32ValueOrDefault(section, key string, defaultVal int32) int32 {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}
	res, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultVal
	}
	return int32(res)
}

func (cfg *ConfigFile) GetFloat32ValueOrDefault(section, key string, defaultVal float32) float32 {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}

	res, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultVal
	}
	return float32(res)
}

func (cfg *ConfigFile) GetFloat64ValueOrDefault(section, key string, defaultVal float64) float64 {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultVal
	}
	return res
}

func (cfg *ConfigFile) GetBoolValueOrDefault(section, key string, defaultVal bool) bool {
	value, err := cfg.getValue(section, key)
	if err != nil {
		return defaultVal
	}

	res, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}
	return res
}

func (cfg *ConfigFile) Reload() {
	cfg.load(cfg.path)
}