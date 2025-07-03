package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileId := "./operation.txt"
	msg := "Friday"

	file, err := os.Create(fileId)
	createPanic(err)

	// the reason for using = instead of := is we are declaring err again

	_, err = io.WriteString(file, msg)
	createPanic(err)

	defer file.Close()

	fmt.Println("ByteData:- ", ReadFile(file.Name()))
}

func createPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(fileId string) string {
	byteData, err := os.ReadFile(fileId)
	createPanic(err)

	return string(byteData)
}
