package main

import "fmt"

func main() {
	str := "Hello 월드!"

	for _, v := range str {
		fmt.Printf("Type: %T Value: %d Char: %c\n", v, v, v)
	}
}
