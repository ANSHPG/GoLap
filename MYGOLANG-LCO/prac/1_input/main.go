package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Lap Time:- ")

	reader := bufio.NewReader(os.Stdin)
	lapTime, _ := reader.ReadString('\n')

	fmt.Printf("Lap time:- %s [recorded]", lapTime)
}
