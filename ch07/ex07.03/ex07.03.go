package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3

	fmt.Println(name, "평균 점수 :", avg)
}

func main() {
	PrintAvgScore("First", 80, 74, 95)
	PrintAvgScore("Second", 88, 92, 53)
	PrintAvgScore("Third", 78, 73, 78)
}
