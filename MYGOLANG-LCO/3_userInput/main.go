package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	//  go to pkg.go.dev to search docu. for pkg(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the time for lap 1")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Printf("Lap time: %s", input)
	fmt.Printf("type: %T", input)
}
