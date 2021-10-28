package main

import (
	"errors"
	"fmt"
)

type myErrorString struct {
	s string
}

func (mError myErrorString) Error() string {
	return mError.s
}

func New(text string) error {
	return myErrorString{text}
}

func main() {
	Error1 := New("error")
	Error2 := New("error")

	if errors.Is(Error1, Error2) {
		fmt.Printf("%v could not equal %v\n", Error1, Error2)
	}
}
