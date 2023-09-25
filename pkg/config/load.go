package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(filename string, cfg any) error {

	y, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer y.Close()

	d := yaml.NewDecoder(y)
	err = d.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

// TODO: go get github.com/spf13/viper 以后换
func LoadChildNode(filename string, cfg any, param ...string) error {

	return nil
	//return c.d.unmarshal(dest, c.cp, param...)
}

func Store(filename string, cfg any) error {
	y, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //os.Create(filename)
	if err != nil {
		return err
	}
	defer y.Close()

	e := yaml.NewEncoder(y)
	defer e.Close()

	err = e.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
