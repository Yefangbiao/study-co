package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

func GetErrorf() error {
	return errors.Errorf("errorf error")
}

func GetErrorNew() error {
	return errors.New("new error")
}

func ErrorWithMessage() error {
	err := GetErrorNew()
	return errors.WithMessage(err, "message: none")
}

func ErrorWrap() error {
	_, err := json.Marshal(make(chan int))
	return errors.Wrap(err, "wrap: none")
}

func main() {
	// 打印堆栈信息
	fmt.Println("-----------------打印堆栈信息-----------------")
	fmt.Printf("%+v\n", ErrorWrap())
	// 错误判断使用 errors.Is 判断
	fmt.Println("-----------------错误判断-----------------")
	error1 := GetErrorf()
	if errors.Is(error1, io.EOF) {
		fmt.Println(error1)
	}
	// 错误类型使用 errors.As 判断
	fmt.Println("-----------------错误类型-----------------")
	error2 := GetErrorNew()
	anotherError := errors.New("another new error")
	if errors.As(error2, &anotherError) {
		// 这里都是内置的error类型，所以返回true
		fmt.Println(error2)
	}
}
