// +build wireinject

package field_struct

import (
	"github.com/google/wire"
)

type Point struct {
	X int
	Y int
}

func NewPointExample() Point {
	return Point{
		X: 5,
		Y: 10,
	}
}

func InitPoint() int {
	panic(wire.Build(
		NewPointExample,
		wire.FieldsOf(new(Point), "X"),
	),
	)
	return 0
}
