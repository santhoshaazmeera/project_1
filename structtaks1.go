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
	name   string
	age    int
	salary string
}

func main() {
	e := employe{name: "dharani", age: 22, salary: 350000.0}
	e1 := employe{name: "adhi", age: 23, salary: 350000.0}
	e2 := employe{name: "sandhya", age: 23, salary: 350000.0}
	e3 := employe{name: "rakshitha", age: 23, salary: 350000.0}
	e4 := employe{name: "sharath", age: 25, salary: 350000.0}

	t := teacher{name: "vignan", age: 28, salary: 800000.0}
	t1 := teacher{name: "narasimham", age: 58, salary: 900000.0}
	t2 := teacher{name: "sudhakar", age: 28, salary: 600000.0}
	s := student{name: "santhosha", age: 21, salary: "no salary fo students"}
	s1 := student{name: "venkky", age: 23, salary: "no salary fo students"}

	implementingmethods(e)
	implementingmethods(e1)
	implementingmethods(e2)
	implementingmethods(e3)
	implementingmethods(e4)
	implementingmethods(t1)
	implementingmethods(t2)
	implementingmethods(s1)
	implementingmethods(t)
	implementingmethods(s)
}

func implementingmethods(u user) {
	u.details()
	u.salaryy()
}
func (e employe) details() {
	fmt.Printf("printing the employee details \n\n" )
	fmt.Printf("name  :%s    age is : %d \n", e.name, e.age)
}
func (s student) details() {
	fmt.Printf("student details \n\n ")
	fmt.Printf("name  :%s and   is %d\n", s.name, s.age)
}
func (t teacher) details() {
	fmt.Printf("teacher details \n \n ")
	fmt.Printf("name  :%s   age  :%d \n", t.name, t.age)
}
func (e employe) salaryy() {
	fmt.Printf("employee %s salary details %f \n", e.name, e.salary)
}

func (t teacher) salaryy() {
	fmt.Printf("teacher %s salary details %f\n\n", t.name, t.salary)
}
func (s student) salaryy() {
	fmt.Println(s.salary, "\n")
}
