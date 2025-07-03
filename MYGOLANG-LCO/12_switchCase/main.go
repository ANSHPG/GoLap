package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Lanes Not Positions!")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	algorithm := r.Intn(6) + 1
	fmt.Println(algorithm)
	//  unlike other language we dont add break here, else we add fallthrough to excite it!
	switch algorithm {
	case 1:
		fmt.Println("Dice val is 1 and yo can open")
	case 2:
		fmt.Println("You can move 2 spots!")
	case 3:
		fmt.Println("You can move 3 spots!")
	case 4:
		fmt.Println("You can move 4 spots!")
	case 5:
		fmt.Println("You can move 5 spots!")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spots and roll dice again!")
	default:
		fmt.Println("What was that?")
	}
}
