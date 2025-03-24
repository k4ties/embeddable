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

//go:embed yaml_test.yaml
var y []byte

func main() {
	conf, err := embeddable.Extract[config](embeddable.YAML, y)
	if err != nil {
		panic(err)
	}
	c := conf.Config

	format := "foo: %s, bar: %d"
	fmt.Printf(format, c.Foo, c.Bar)
}
