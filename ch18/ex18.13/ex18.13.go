package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := []Student{
		{"Jang", 24},
		{"Kim", 22},
		{"Lee", 23},
		{"Park", 20},
		{"Song", 21},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
}
