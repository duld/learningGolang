/*
This exercise will reinforce our understanding of method sets:

create a type person struct
attach a method speak to type person using a pointer receiver
*person
create a type human interface
to implicitly implement the interface, a human must have the speak method
create func “saySomething”
have it take in a human as a parameter
have it call the speak method
show the following in your code
you CAN pass a value of type *person into saySomething
you CANNOT pass a value of type person into saySomething
*/
package main

import "fmt"

type Person struct {
	First, Last string
	Age         int
}

func (p *Person) Speak() {
	fmt.Printf("Hi, my name is %v %v, nice to meet you\n", p.First, p.Last)
}

type Human interface {
	Speak()
}

func saySomething(h Human) {
	h.Speak()
}

func main() {
	p1 := Person{
		First: "J",
		Last:  "M",
		Age:   32,
	}

	saySomething(&p1) // runs
	//saySomething(p1) // wont run
}
