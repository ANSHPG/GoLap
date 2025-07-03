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
	triggerPanic(err)
	defer file.Close()

	// the reason for using = instead of := is we are declaring err again

	_, err = io.WriteString(file, msg)
	triggerPanic(err)

	fmt.Println("ByteData:- ", ReadFile(file.Name()))
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(fileId string) string {
	byteData, err := os.ReadFile(fileId)
	triggerPanic(err)

	return string(byteData)
}
