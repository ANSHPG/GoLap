package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("recieving data from AMG servers")
	PerformGetRequest()
}

func PerformGetRequest() {
	const driverUrl string = "https://f1api.dev/api/drivers/search?q=verstappen"

	response, err := http.Get(driverUrl)
	triggerPanic(err)

	defer response.Body.Close()

	fmt.Println("status: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)
	/*
		if the content length is  -1 =>
		1- The server did not send a Content-Length header.
		2- The response is chunked (i.e., sent in pieces without a predefined length).
	*/

	/*
		Imagine =>
			You're writing words on a whiteboard — instead of getting a new whiteboard for every word (which is slow and wasteful), you keep writing on the same one. That’s what strings.Builder helps you do!

			It's a special tool in Go that helps you build strings piece by piece efficiently — without creating a new string every time you add something.

			basically helps in memory allocation and management!

	*/
	var responseString strings.Builder

	/*
		response.body =>
		You're calling a restaurant and asking for their menu. After some time, they start reading it out slowly, line by line over the phone.

		That live stream of words from the other end?
		That's like response.Body.

		It's a stream — not the full content at once — it's like reading from a file or pipe.

		content, err := io.ReadAll(response.Body)

	*/
	content, err2 := io.ReadAll(response.Body)
	triggerPanic(err2)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
