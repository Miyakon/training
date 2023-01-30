package main

import "fmt"

type Person struct {
	Name     string
	Favorite string
}

func main() {
	fmt.Println("Hello")

	pointerVal := new(Person)
	pointerVal.Name = "rovelte"
	pointerVal.Favorite = "cooking"

	fmt.Printf("構造体 = %p\n", pointerVal)

	i := 42

	p := &i
	fmt.Println(*p)

	pointerVal.Greet("my name")
	println(pointerVal.Name)
}

func (p *Person) Greet(msg string){
	(*p).Name = msg
	println(p.Name)
}
