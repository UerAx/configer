package configer

import "testing"

func TestReader(t *testing.T) {
	cfg, err := LoadConfigFile("./config.conf")
	if err != nil {
		t.Fatal(err)
	}
	if 1 != cfg.GetIntValueOrDefault("test", "test", 1) {
		t.Fatal()
	}
	if 1 != cfg.GetIntValueOrDefault("test", "test5", 1) {
		t.Fatal()
	}
	if cfg.GetBoolValueOrDefault("test", "test2", false) {
		t.Fatal()
	}
	cfg.Reload()
	if 1 != cfg.GetIntValueOrDefault("test", "test", 1) {
		t.Fatal()
	}

}
