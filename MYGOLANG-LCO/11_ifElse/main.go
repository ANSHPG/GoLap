package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("No If NO Else Only Pace!")

	fmt.Printf("Enter min:- ")
	readMin := bufio.NewReader(os.Stdin)
	inputMin, _ := readMin.ReadString('\n')
	min, _ := strconv.ParseInt(strings.TrimSpace(inputMin), 10, 32)

	fmt.Printf("Enter sec:- ")
	readSec := bufio.NewReader(os.Stdin)
	inputSec, _ := readSec.ReadString('\n')
	sec, _ := strconv.ParseInt(strings.TrimSpace(inputSec), 10, 32)

	lapTime := (min * 60) + sec
	var status string

	if lapTime < 73 {
		status = "Fast"
	} else if lapTime < 83 {
		status = "Average"
	} else {
		status = "Slow"
	}

	fmt.Printf("\n---Status---\nLapTime:- %d sec(s)\nReaction:- %s\n\n", lapTime, status)

	if num := 19; num < int(lapTime) {
		fmt.Println("num is fast!")
	} else {
		fmt.Println("num is slow!")
	}

	// if err != nil {

	// }
}
