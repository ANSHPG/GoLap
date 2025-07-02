package main

import "fmt"

/* Arrays are not used much in goLang, in GO we use Slices */
func main() {
	fmt.Println("A Ray of Lightning!")

	var franchise [4]string

	franchise[0] = "McLaren"
	franchise[1] = "AstonMartin"
	franchise[3] = "LaFerrari"

	fmt.Println("Top Teams:- ", franchise)
	fmt.Println("Length:- ", len(franchise))

	var location = [3]string{"Bahrain", "Monaco", "SaudiArabia"}
	fmt.Println("Locations:-", location)

}
