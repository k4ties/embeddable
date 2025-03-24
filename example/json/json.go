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

//go:embed json_test.json
var j []byte

func main() {
	conf, err := embeddable.Extract[config](embeddable.JSON, j)
	if err != nil {
		panic(err)
	}
	c := conf.Config

	format := "foo: %s, bar: %d"
	fmt.Printf(format, c.Foo, c.Bar)
}
