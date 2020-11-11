package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Reader interface {
	ReadConf() (Config, error)
}

type FileReader struct {
	file string
}

func NewFileReader(file string) FileReader {
	return FileReader{
		file: file,
	}
}

func (f FileReader) ReadConf() (Config, error) {
	content, err := ioutil.ReadFile(f.file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
