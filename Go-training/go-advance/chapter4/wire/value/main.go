package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	point := InitPoint()
	fmt.Println(point.X)
	fmt.Println(point.Y)
}
