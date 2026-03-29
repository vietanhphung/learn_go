package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[1] = 1
	a[2] = 2
	b := [3]int{1}
	c := [...]int{1, 2}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("Length of a is : ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(i, "th element of a is: ", a[i])
	}

	for i, v := range a {
		fmt.Println(i, "th element of a is: ", v)
	}

	matrixA := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println("original matrixA: ", matrixA)

	slicedA := matrixA[1:3]
	fmt.Println("sliced matrix A is : ", slicedA)

	// modify each element of slice, this share the same memory (reference to memory, not creating new variable), hence modify the original matrix, use copy() to make a new matrix
	slicedA[0] = [2]string{"duck", "pig"}
	slicedA[1] = [2]string{"pigeon", "peacock"}

	fmt.Println("modified matrix A: ", matrixA)

	// capacity of a slice
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing fruitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	// Slice(dynamic) is reference to array(fixed length)
	// sliceA = append(type, value) -> original array length +1 and capacity is double
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	// memory optimize: since slice references array, as long as we use slice, array (can be very large) cannot be garbage collected
	// use copy() if we only need a small section of the array, then the array can be garbage collected
}
