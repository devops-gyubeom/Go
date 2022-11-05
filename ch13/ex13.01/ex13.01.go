package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "Seoul"
	house.Size = 28
	house.Price = 9.8
	house.Type = "Apart"

	fmt.Println("Address:", house.Address)
	fmt.Printf("Size: %dm^2\n", house.Size)
	fmt.Printf("Price: %.2f won\n", house.Price)
	fmt.Println("Type:", house.Type)
}
