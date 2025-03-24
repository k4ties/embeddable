package main

import (
	_ "embed"
	"fmt"
	"github.com/k4ties/embeddable"
)

type config struct {
	Config struct {
		Foo string
		Bar int
	}
}

//go:embed toml_test.toml
var t []byte

func main() {
	conf, err := embeddable.Extract[config](embeddable.TOML, t)
	if err != nil {
		panic(err)
	}
	c := conf.Config

	format := "foo: %s, bar: %d"
	fmt.Printf(format, c.Foo, c.Bar)
}
