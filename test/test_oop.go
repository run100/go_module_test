package main

import "fmt"

type Person struct {
	name string
}

type Women struct {
	Person
	sex string
}

func main() {
	women := Women{Person{"zzzz"}, "111"}
	fmt.Printf("women %v", women)
}
