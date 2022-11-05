package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p value: %p\n", p)
	fmt.Printf("p memory value: %d\n", *p)

	*p = 100

	fmt.Printf("a value: %d\n", a)
}
