package minecraft

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Server struct {
	Dir     string `json:dir`
	JarPath string `json:jar_path`
	Name    string `json:name`
}

type Config struct {
	Servers []Server `json:servers`
}

func filename() string {
	home, _ := os.LookupEnv("HOME")
	return filepath.Join(home, ".gsmconfig")
}

func (cfg *Config) Save() error {
	var filename = filename()
	return cfg.SaveFile(filename)
}

func (cfg *Config) SaveFile(filename string) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, jsonData, 0600)
	return err
}

func LoadConfig() (Config, error) {
	var filename = filename()
	return LoadConfigFile(filename)
}

func LoadConfigFile(filename string) (Config, error) {
	var cfg Config
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		cfg.Save()
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal([]byte(file), &cfg)
	return cfg, err
}
