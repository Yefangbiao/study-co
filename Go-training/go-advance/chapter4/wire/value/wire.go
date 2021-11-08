// +build wireinject

package main

import "github.com/google/wire"

func InitPoint() *Point {
	panic(wire.Build(wire.Value(&Point{
		X: 1,
		Y: 2,
	})))
	return &Point{}
}
