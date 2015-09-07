package main

import (
	"fmt"
	"regexp"

	"github.com/naoina/toml"
)

type Config struct {
	Lints []Lint
}

func (c *Config) FindLint(name string) (*Lint, error) {
	for _, l := range c.Lints {
		if l.Name == name {
			return &l, nil
		}
	}
	return nil, fmt.Errorf("%s does not found", name)

}

type Regexp regexp.Regexp

// TODO: 正規表現以外
type Lint struct {
	Name    string  `toml:"name"`
	Regexp  *Regexp `toml:"regexp"`
	FName   int     `toml:"fname"`
	Line    int     `toml:"line"`
	Col     int     `toml:"col"`
	Message int     `toml:"message"`
}

func (re *Regexp) UnmarshalTOML(b []byte) error {
	r, err := regexp.Compile(string(b))
	if err != nil {
		return err
	}
	*re = Regexp(*r)
	return nil
}

var _ toml.Unmarshaler = &Regexp{}

var config *Config = func() *Config {
	b, err := Asset("conf/default.toml")
	if err != nil {
		panic(err)
	}

	conf := &Config{}
	err = toml.Unmarshal(b, conf)
	if err != nil {
		panic(err)
	}

	return conf
}()
