package minecraft

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestConfig_SaveAndLoad(t *testing.T) {
	cfg := Config{
		Servers: []Server{
			{JarPath: "/path/to/minecraft/server1.jar",
				Name: "server1",
				Dir:  "/path/to/server1"},
			{JarPath: "/path/to/minecraft/server2.jar",
				Name: "server2",
				Dir:  "/path/to/server2"},
			{JarPath: "/path/to/minecraft/server3.jar",
				Name: "server3",
				Dir:  "/path/to/server3"},
		},
	}
	tempFile, err := ioutil.TempFile("", "testconfig")
	if err != nil {
		t.Errorf("Config tempfile create failed: %v", err)
	}
	filename := tempFile.Name()
	err = cfg.SaveFile(filename)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}
	cfg1, err := LoadConfigFile(filename)
	if err != nil {
		t.Errorf("Config load failed: %v", err)
	}
	if len(cfg.Servers) != len(cfg1.Servers) {
		t.Errorf("len(servers) does not match after round trip %v != %v",  len(cfg.Servers), len(cfg1.Servers) )
	}
	cfgbytes, _ := json.Marshal(cfg)
	cfg1bytes, _ := json.Marshal(cfg1)
	if string(cfgbytes) != string(cfg1bytes) {
		t.Errorf("start config string != loaded config string: %v != %v",
			string(cfgbytes),
			string(cfg1bytes))
	}
}
