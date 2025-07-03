package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://f1api.dev/api/drivers/search?q=lewis"

func main() {
	fmt.Println("Request to Box, Request to Lose!")

	response, err := http.Get(url)
	triggerPanic(err)

	fmt.Printf("%T\n", response)
	// fmt.Println(response)

	databytes, err := io.ReadAll(response.Body)
	triggerPanic(err)

	content := string(databytes)
	fmt.Println(content)

	defer response.Body.Close() // Caller's Responsibility to close the request
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
