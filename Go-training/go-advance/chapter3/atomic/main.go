package main

import (
	"fmt"
	"sync/atomic"
)

type Coin struct {
	Num   int
	Total int
}

func main() {
	c1 := &Coin{
		Num:   5,
		Total: 15,
	}
	c2 := &Coin{
		Num:   10,
		Total: 20,
	}

	var value atomic.Value

	var loop0, loop1 func()
	loop0 = func() {
		value.Store(c1)
		go loop1()
	}

	loop1 = func() {
		value.Store(c2)
		go loop0()
	}

	go loop0()
	for {
		fmt.Println(value.Load())
	}
}
