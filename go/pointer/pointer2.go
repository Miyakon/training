package main

import "fmt"

func main() {
	var x = 100
	fmt.Println("x address:\t", &x)
    // xアドレス

	var y *int
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)
    // y値
    // yアドレス

	y = &x
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)
    // xアドレス
    // yアドレス

    fmt.Print("\n\n関数引き渡し\n")

	var alice = "alice"
    fmt.Println("\nalice address:\t", &alice)
    b := &alice
    fmt.Println("b value:\t", b)
    fmt.Println("b address:\t", &b)
    fmt.Println("b refer:\t", *b)

    // alice変数のポインターbを渡す
	show(b)
    fmt.Println("b refer:\t", *b)  // bobのつもりがaliceのまま
}

// sに渡されるのはstringのポインター
func show(s *string) {
    tmp := "bob"
    fmt.Println("tmp address:\t", &tmp)

    fmt.Println("s value:\t", s)
    fmt.Println("s address:\t", &s)
    // 受け取った値は

    s = &tmp  // 上書きのつもり
    fmt.Println("s value:\t", s)
    fmt.Println("s address:\t", &s)
}