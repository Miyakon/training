package main

import "fmt"

type Square struct {
	X, Y int
}

func (s *Square) Area() int {
	return s.X * s.Y
}

func main() {
	s := Square {5, 5}
	fmt.Println(s.Area())
}

