package main

import (
	"fmt"
)

type MyError struct {
	Msg  string
	File string
	Line int
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", m.File, m.Line, m.Msg)
}

func New() error {
	return &MyError{
		Msg:  "File can not found",
		File: "test.txt",
		Line: 32,
	}
}

func main() {
	err:=New()

	switch err.(type) {
	case *MyError:
		fmt.Println("error:", err)
	}
}
