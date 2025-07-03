package main

import "fmt"

func main() {
	fmt.Println("Defer the Odds!")

	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("World")
	letsDefer()
}

// World
// letsDefer() => 3,2,1
// LIFO OF {defer} :- Two, One, Hello

func letsDefer() {
	for p := 1; p < 4; p++ {
		defer fmt.Println(p)
	}
}
