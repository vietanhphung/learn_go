package main

import (
	"fmt"
	"math"
)

func main() {
	helloWorld()
	myAge()
	compare()
	constCheck()
	fmt.Println(calculateBill(5, 10, 10))
}

func helloWorld() {
	fmt.Println("Hello World")
}

func myAge() {
	var age int
	age = 10
	var catAge int = age * 7
	shortHandVar := "short hand"
	fmt.Println("My age is", age)
	fmt.Println("My cat's age is", catAge)
	fmt.Println("Testing short hand initializaion", shortHandVar)
}

func compare() {
	a, b := 145.5, 32.2
	c := math.Min(a, b)
	fmt.Println("min is: ", c)
}

func constCheck() {
	const pi = 3.1415926
	fmt.Println(pi)
}

// --- func fucntionName(paremeterName1 datatype1, parameterName2 dateType2) returnType {}

func calculateBill(price int, quantity int, tips int) (int, int) {
	var totalPrice = price * quantity
	totalPrice = totalPrice + totalPrice*tips/100
	return totalPrice, tips
}
