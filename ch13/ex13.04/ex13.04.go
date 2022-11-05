package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"Kim", "kim", 22, 10}
	vip := VIPUser{
		User{"Jang", "jang", 27, 10},
		250,
		3,
	}

	fmt.Printf("User: %s ID: %s Age: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP User: %s ID: %s Age: %d VIP Level: %d User Level: %d\n", vip.Name, vip.ID, vip.Age, vip.Level, vip.User.Level)
}
