package main

import "fmt"

type Square struct {
	x, y int
}

func Area(s Square) int {
	return s.x * s.y
}

func main() {
	s := Square {5, 5}
	fmt.Println(Area(s))
}