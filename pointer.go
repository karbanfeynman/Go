package main

import "fmt"

type Person struct {
	age int
}

func (p *Person)getAge() int {
	return p.age
}

func (p *Person)setAge(age int) {
	p.age = age
}

func main() {
	p := new(Person)

	p.age = 22
	fmt.Println(p.age)
	p.setAge(10)
	fmt.Printf("Age: %d\n", p.getAge())
}
