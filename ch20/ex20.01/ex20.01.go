package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Hi! I'm %s and %d years old", s.Name, s.Age)
}

func main() {
	student := Student{"Kim", 12}
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())
}
