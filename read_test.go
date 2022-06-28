package configer

import "testing"

func TestReader(t *testing.T) {
	cfg, err := LoadConfigFile("./config.conf")
	if err != nil {
		t.Fatal(err)
	}
	print(cfg.GetStringValueOrDefault("", "bbb","1"))
	cfg.Reload()
	print(cfg.GetStringValueOrDefault("", "bbb","1"))


}
