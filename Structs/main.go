package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{"blob", 20})
	fmt.Println(Person{name: "alice", age: 30})
	fmt.Println(Person{name: "fred"})
	fmt.Println(&Person{name: "ann", age: 40})
	fmt.Println(newPerson("jon"))

	s := Person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
