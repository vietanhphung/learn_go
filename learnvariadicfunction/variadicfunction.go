package main

import (
	"fmt"
)

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	//passing a slice to a variadic function -> add ... suffix
	nums := []int{89, 90, 95}
	find(89, nums...)

}

func find(num int, nums ...int) {
	fmt.Println("Type of nums: %T", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println("Found num: ", v, " at index ", i)
			found = true
		}
	}
	if !found {
		fmt.Println(" num: ", num, " not found")
	}
}

// variadic arguments uses slide, so the equivalent is

func findWSlide(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

// but variadic is more readable and doesn't need to cast type each time call function
// find(89, []int{89, 90, 95})
