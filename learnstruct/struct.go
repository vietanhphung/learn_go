package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	emp1 := Employee{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}
	fmt.Printf("Employee: %+v\n", emp1)

	emp2 := Employee{"Jane", "Smith", 25}
	fmt.Printf("Employee: %+v\n", emp2)

	//anonymous struct
	emp3 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Alice",
		lastName:  "Johnson",
		age:       28,
	}
	fmt.Printf("Employee: %+v\n", emp3)

	// Accessing struct fields
	fmt.Printf("First Name: %s\n", emp1.firstName)
	fmt.Printf("Last Name: %s\n", emp1.lastName)
	fmt.Printf("Age: %d\n", emp1.age)

	// Modifying struct fields
	emp1.age = 31
	fmt.Printf("Updated Age: %d\n", emp1.age)

	// Structs can be compared for equality if all their fields are comparable
	emp4 := Employee{"John", "Doe", 31}
	fmt.Printf("emp1 == emp4: %t\n", emp1 == emp4) // true

	// Structs can be used as map keys if all their fields are comparable
	employeeMap := make(map[Employee]string)
	employeeMap[emp1] = "Employee 1"
	employeeMap[emp2] = "Employee 2"
	fmt.Println(employeeMap)

	//zero value of struct
	var emp5 Employee
	fmt.Printf("Zero value: %+v\n", emp5)

	// Struct can be specfied with only some fields, the rest will be zero value
	emp6 := Employee{firstName: "Bob"}
	fmt.Printf("Employee: %+v\n", emp6)

	//Pointer to struct
	emp7 := &Employee{
		firstName: "Charlie",
		lastName:  "Brown",
		age:       35,
	}
	fmt.Println("Employee (by pointer): ", (*emp7))

	//Anoymous field struct -> no name so the type is used as the field name, also called embedding
	type Person struct {
		string
		int
	}

	p1 := Person{"David", 40}
	fmt.Println("Person: ", p1)
	fmt.Println("Person Name (access by type name): ", p1.string)

	// Nested struct
	type Company struct {
		name     string
		employee Employee
	}

	company := Company{
		name: "Tech Co",
		employee: Employee{
			firstName: "Eve",
			lastName:  "Adams",
			age:       29,
		},
	}
	fmt.Printf("Company: %+v\n", company)

	// ** Promoted fields in embedded struct
	type Manager struct {
		Employee   // embedding Employee struct
		department string
	}

	m1 := Manager{
		Employee: Employee{
			firstName: "Frank",
			lastName:  "Miller",
			age:       45,
		},
		department: "Sales",
	}
	fmt.Printf("Manager: %+v\n", m1)
	fmt.Printf("Manager's First Name (promoted field): %s\n", m1.firstName) // Accessing promoted field directly

	//If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages. Similarly, if the fields of a struct start with caps, they can be accessed from other packages

	// Struct variables are not comparable if they contain fields that are not comparable
	/*
		type image struct {
			data map[int]int
		}

		image1 := image{
			data: map[int]int{
				0: 155,
			}}
		image2 := image{
			data: map[int]int{
				0: 155,
			}}

			if image1 == image2 { // compile error: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
				fmt.Println("image1 and image2 are equal")
			}
	*/
}
