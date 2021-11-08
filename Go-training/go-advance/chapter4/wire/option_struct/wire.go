// +build wireinject

package main

import (
	"github.com/google/wire"
)

type typeA int
type typeB int

type Options struct {
	A typeA
	B typeB
}

type Greeter struct {
	AAA typeA
	BBB typeB
}

func NewGreeter(opts *Options) *Greeter {
	// ...
	return &Greeter{
		AAA: opts.A,
		BBB: opts.B,
	}
}

var GreeterSet = wire.NewSet(wire.Struct(new(Options), "*"), NewGreeter)

func InitGreeter() *Greeter {
	panic(wire.Build(
		wire.Value(typeA(10)),
		wire.Value(typeB(20)),
		GreeterSet,
	),
	)
	return &Greeter{}
}
