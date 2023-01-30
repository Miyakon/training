
package main

import "fmt"

type Square struct {
	X, Y int
}

func (s *Square) Area() int {
	return s.X * s.Y
}

func (s *Square) Change() {
	s.X = 10
	s.Y = 10
}

func main() {
	s := Square {5, 5}
	fmt.Println(s.Area())
	s.Change()
	fmt.Println(s.Area())
}