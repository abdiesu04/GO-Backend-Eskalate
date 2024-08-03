package main

import (
	"fmt"
)

type Employee interface{
	UpdateName() 
	PaySalary()
	GiveBadgePoint() 
}

type Student struct{
	Name string
	Salary int
	BadgePoint int
	SalaryStatus bool
}

func (s *Student) UpdateName(name string){
	s.Name = name
}

func (s * Student) PaySalary(){
	s.SalaryStatus = true
}

func (s *Student) GiveBadgePoint() {
 	s.BadgePoint += 100
}
func main() {
	fmt.Println("Hello, World!")
	/*
	type Student struct{
	Name string
	Salary int
	BadgePoint int
	SalaryStatus bool
}*/

	student := Student{"John", 1000, 100, false}
fmt.Println(student)	

	student.UpdateName("Doe")
	student.PaySalary()
	student.GiveBadgePoint()

	fmt.Println(student)


}

