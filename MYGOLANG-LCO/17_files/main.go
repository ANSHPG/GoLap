package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Acess the MotorSport Files!")

	content := "The Final is at Abu Dhabi"

	file, err := os.Create("./motorsport.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()

	readFile("./motorsport.txt")
}

func readFile(flName string) {
	dataByte, err := os.ReadFile(flName)

	checkNilErr(err)

	fmt.Println("Text Data inside file:- ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
