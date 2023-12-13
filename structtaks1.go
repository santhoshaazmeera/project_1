package main

import (
	"fmt"
)

type user interface {
	details()
	salaryy()
}
type employe struct {
	name   string
	age    int
	salary float64
}
type teacher struct {
	name   string
	age    int
	salary float64
}
type student struct {
	name string
	age  int
}

func main() {
	e := employe{name: "dharani", age: 22, salary: 350000.0}
	t := teacher{name: "vignan", age: 28, salary: 800000.0}
	s := student{name: "santhosha", age: 21}
	implementingmethods(e)
	implementingmethods(t)
	implementingmethods(s)
}

func implementingmethods(u user) {
	u.details()
	u.salaryy()
}
func (e employe) details() {
	fmt.Println("printing the eploye details ", e.name, e.age)
}
func (s student) details() {
	fmt.Println("student details are ", s.name, s.age)
}
func (t teacher) details() {
	fmt.Println("teacher details are ", t.name, t.age)
}
func (e employe) salaryy() {
	fmt.Println("printing the eploye details ", e.salary)
}

func (t teacher) salaryy() {
	fmt.Println("teacher details are ", t.salary)
}
