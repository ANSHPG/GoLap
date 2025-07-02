package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Lap :- ")
	lap, _ := reader.ReadString('\n')
	lapInt, err := strconv.ParseUint(strings.TrimSpace(lap), 10, 32)
	// uint, 32 -> 0 - 65535

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Enter Lap Time:- ")
		lapTime, _ := reader.ReadString('\n')
		timeInt, err1 := strconv.ParseUint(strings.TrimSpace(lapTime), 10, 64)

		if err1 != nil {
			fmt.Println(err1)
		} else {
			min := timeInt / 60
			sec := timeInt % 60
			fmt.Printf("\n---Report---\nLap:- %d\nLap Time:- %dmin %dsec", lapInt, min, sec)
		}
	}

}
