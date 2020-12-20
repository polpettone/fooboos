package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func loadFooboos(file string) (*Fooboos, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fooboos := &Fooboos{}
	err = yaml.Unmarshal(content, &fooboos)
	if err != nil {
		return nil, err
	}
	return fooboos, nil
}

func loadRaw(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func writeRaw(file string, content []byte) error {
	err := ioutil.WriteFile(file, content, 744)
	return err
}
