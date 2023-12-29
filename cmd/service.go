package cmd

import (
	"os"

	"gopkg.in/yaml.v2"
)

func loadFooboos(file string) (*Fooboos, error) {
	content, err := os.ReadFile(file)
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
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func writeRaw(file string, content []byte) error {
	err := os.WriteFile(file, content, 744)
	return err
}
