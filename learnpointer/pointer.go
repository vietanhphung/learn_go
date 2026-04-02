package main

import (
	"fmt"
)

func change(val *int) { //passing pointer as argument
	*val = 55
}

func returnPointer() *int { // return pointer
	i := 5
	return &i
}

func modifyWithPointer(arr *[3]int) { //pointer to array
	(*arr)[0] = 90
	//or equivalently, arr[0]
}

func modifyWithSlice(arr []int) {
	arr[0] = 90
}

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is ", a)

	// zero value of pointer is nil

	// new() takes a type as argument and returns a pointer
	// to newly allocated zero value of the type passed
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85 //assign value 85 to the memory location
	fmt.Println("New size value is ", *size)

	// dereferencing a pointer
	b0 := 255
	a0 := &b0
	fmt.Println("address of b0 is", a0)
	fmt.Println("value of b0 is", *a0)
	*a0++
	fmt.Println("new value of b0 is", b0)

	// passing pointer as argument
	change(a0)
	fmt.Println("value of b0 is", *a0)

	d := returnPointer()
	fmt.Println("a function return pointer, value of d is", *d)

	//Even though it works, we should pass slice to modify an array instead of pointer to array
	a1 := [3]int{89, 90, 91}
	fmt.Println("value of a1 is", a1)

	modifyWithPointer(&a1)
	fmt.Println("modified with Pointer: ", a1)

	a1 = [3]int{89, 90, 91}
	modifyWithSlice(a1[:])
	fmt.Println("modified with Slice: ", a1)

	//Unlike C and CPP, pointer in Go does not support arithmetic
	b1 := [...]int{109, 110, 111}
	p := &b1
	fmt.Println("address of b1 is", p, "uncomment the following line will make the program crash as we cannot do arithmetic on pointer")
	//p++

}
