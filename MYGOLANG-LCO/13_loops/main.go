package main

import "fmt"

func main() {
	fmt.Println("Do 78 loops to win Monaco GP!")
	grandPrix := []string{"Monaco", "British", "Belgian", "Italian", "Brazilian", "Singapore", "AbuDhabi"}

	//  no ++d prefix
	for location := 0; location > len(grandPrix); location++ {
		fmt.Println("<- ", grandPrix[location], " ->")
	}

	// for location := range grandPrix {
	// 	fmt.Println("<- ", grandPrix[location], "Grand Pix\t->")
	// }

	// for _, val := range grandPrix {
	// 	// fmt.Printf("location:- %v\tGP:- %v\n", index, val)
	// 	fmt.Printf("GP:- %v\n", val)
	// }

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at Lap 22")
}
