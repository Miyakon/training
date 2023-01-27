package main

import "fmt"

func main() {
	var x = 100
	fmt.Println("x address:\t", &x)

	var y *int
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)
	y = &x
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)

	var z **int
    fmt.Println("z value:\t", z)
    fmt.Println("z address:\t", &z)
    z = &y
    fmt.Println("z value:\t", z)
    fmt.Println("z address:\t", &z)

	var a = "alice"
    fmt.Println("\na address:\t", &a)
    b := &a
    fmt.Println("b value:\t", b)
    fmt.Println("b address:\t", &b)

	show(b)
    fmt.Println(*b)  // bobのつもりがaliceのまま
}

func show(s *string) {
    tmp := "bob"
    fmt.Println("tmp address:\t", &tmp)

    fmt.Println("s value:\t", s)
    fmt.Println("s address:\t", &s)

    *s = tmp  // 上書きのつもり
    fmt.Println("s value:\t", s)
    fmt.Println("s address:\t", &s)
}