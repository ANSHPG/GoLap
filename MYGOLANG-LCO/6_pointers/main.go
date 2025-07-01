package main

import "fmt"

func main() {
	fmt.Println("Point to curve ahead!")

	/*
		A pointer is like a label that tells you where a value is stored in memory, not the value itself.
		Instead of holding a number or string, it holds the address where that number or string lives.
	*/

	// var ptr *int
	// fmt.Println(ptr)

	num := 75
	var ptr = &num

	fmt.Printf("val:- %d\naddress:- %p", *ptr, ptr)
	*ptr = *ptr * 2
	fmt.Printf("\n%d : %p", *ptr, ptr)
}
