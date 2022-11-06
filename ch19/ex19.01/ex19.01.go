package main

import "fmt"

type account struct {
	balance int
}

// general function
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// method
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}

	withdrawFunc(a, 30)  // call function
	a.withdrawMethod(30) // call method

	fmt.Printf("%d\n", a.balance)
}
