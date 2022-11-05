package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"Kim", "kim", 22}
	vip := VIPUser{
		User{"Jang", "jang", 27},
		3,
		250,
	}

	fmt.Printf("User: %s ID: %s Age: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP User: %s ID: %s Age: %d VIP Level: %d VIP Price: %d won\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price)
}
