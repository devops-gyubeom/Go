package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Printf("Input number : ")

		n, err := InputIntValue()

		if err != nil {
			fmt.Println("Only number please")
		} else {
			if n > r {
				fmt.Println("Your number is greater!")
			} else if n < r {
				fmt.Println("Your number is lesser!")
			} else {
				fmt.Println("Congratulations! Number of attmeps :", cnt)
				break
			}

			cnt++
		}
	}
}
