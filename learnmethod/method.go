package main

// Methods in Go are functions that are associated with a specific type.
// They allow you to define behavior for your types and can be called on instances of those types.
// Methods are defined using a receiver, which is a special parameter that specifies the type the method is associated with.

// func (t Type) methodName(parameter list) {}

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
	age      int
}

/*
displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalaryMethod() {
	fmt.Printf("print by method: Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// function displaySalary() is a standalone function that takes an Employee as a parameter
func displaySalaryFuction(e Employee) {
	fmt.Printf("print by function: Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}

	// method vs function
	// A method is a function that is associated with a specific type, while a function is a standalone block of code that can be called independently.
	// Methods have a receiver, which is the type they are associated with, while functions do not have a receiver.

	emp1.displaySalaryMethod()
	displaySalaryFuction(emp1)

	/*
		Go is not a pure object-oriented programming language and it does not support classes.
		Hence methods on types are a way to achieve behaviour similar to classes.
		Methods allow a logical grouping of behaviour related to a type similar to classes.
		In the above sample program, all behaviours related to the Employee type can be grouped
		by creating methods using Employee receiver type.
		For example, we can add methods like calculatePension, calculateLeaves and so on.

		Methods with the same name can be defined on different types whereas functions with the same names are not allowed.
		Let’s assume that we have a Square and Circle structure.
		It’s possible to define a method named Area on both Square and Circle.
	*/

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	// The difference between VALUE and POINTER receiver is,
	// changes made inside a method with a pointer receiver is visible to the caller
	// whereas this is not the case in value receiver.

	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change (by value receiver): %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change (by value receiver): %s", e.name)

	fmt.Printf("\n\nEmployee age before change (by pointer receiver): %d", e.age)
	(&e).changeAge(51) // (e).changeAge(51) // Go automatically converts value to pointer when calling a method with pointer receiver
	fmt.Printf("\nEmployee age after change (by pointer receiver): %d\n", e.age)

	// when to use value receiver and when to use pointer receiver?
	// 1. If the method needs to modify the receiver, then you should use a pointer receiver.
	// 2. If the method does not need to modify the receiver and the receiver is a small struct, then you can use a value receiver.
	// 3. If the method does not need to modify the receiver and the receiver is a large struct, then you should use a pointer receiver to avoid copying the entire struct.

	/*
		When a function has a value argument, it will accept only a value argument.
		When a function has a pointer argument, it will accept only a pointer argument.
		When a method has a value receiver, it will accept both pointer and value receivers.
		When a method has a pointer receiver, it will accept both pointer and value receivers.
	*/

	/*
		Method with no-struct receiver
		definition must be in the same package
		otherwise create a type alias (for this buit-in type) and define method on that type alias:

		type myInt int

		func (a myInt) add(b myInt) myInt {
			return a + b
		}
	*/

}
