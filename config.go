package main

import (
	"regexp"

	"github.com/naoina/toml"
)

type Config struct {
	Lints []Lint
}

type Regexp regexp.Regexp

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
