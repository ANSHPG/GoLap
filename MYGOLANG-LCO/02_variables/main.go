package main

import "fmt"

const loginToken string = "position:1"

func hide() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var num int16 = -675
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	var floatVal float64 = -675.47447647956
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	var val int
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n", val)

	// implicit type
	var web = "f1.com" // lexar helps here to determine variable
	fmt.Println(web)
	fmt.Printf("type: %T", web)
}

func main() {
	// no var style
	load := 7500 // scope is only inside of method not outside like var
	fmt.Printf("val: %d \ntype: %T", load, load)

	fmt.Printf("\nval: %s \ntype: %T\nstatus: %t", loginToken, loginToken, true)

	// %s -> string
	// %d -> integer numeric
	// %g -> decimal numeric
	// %t -> boolean
	// %T -> type

}
