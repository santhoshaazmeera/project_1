package main

import "fmt"

// import "fmt"
type vehicle interface {
	start()
}
type car struct {
	c string
}
type bicycle struct {
	b string
}

func main() {
	d := car{c: "hello iam "}
	drive(d)
	bi := bicycle{b: "helloo"}
	drive(bi)
}
func drive(v vehicle) {
	v.start()
}
func (s car) start() {
	fmt.Printf("%s starting  car .\n", s.c)
}
func (b bicycle) start() {
	fmt.Printf(" %s starting bicycle ..", b.b)
}




