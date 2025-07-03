package main

import "fmt"

/*
	Defer => Basically like other lngauages GO executes line by line, when we use the keyword defer along side a line, it sends it to end of the execution stack of a function, and during multiple defers it behaves as LIFO so the last defer line exeucutes first in the defer list.
*/

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
