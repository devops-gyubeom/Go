package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func main() {
	for {
		fmt.Printf("Input number : ")

		n, err := InputIntValue()

		if err != nil {
			fmt.Println("Only number please")
		} else {
			fmt.Println("number :", n)
		}
	}
}
