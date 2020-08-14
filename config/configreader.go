package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigReader interface {
	Read() (Config, error)
}

type JsonConfigReader struct {
	fileName string
}

func NewJsonConfigReader(fileName string) ConfigReader {
	return JsonConfigReader{fileName: fileName}
}

func (r JsonConfigReader) Read() (Config, error) {
	var config Config

	ex, err := os.Executable()
	if err != nil {
		return config, err
	}

	exPath := filepath.Dir(ex)
	configContent, _ := ioutil.ReadFile(filepath.Join(exPath, r.fileName))
	err = json.Unmarshal(configContent, &config)
	return config, err
}
