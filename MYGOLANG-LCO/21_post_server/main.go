package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("sending data to AMG servers..")
	PostJsonRequest()
}

func PostJsonRequest() {
	const server string = "https://httpbin.org/post"

	//  fakse json payload

	requestBody := strings.NewReader(`
		{
			"company" : "Mercedes AMG",
			"net_value" : 3.8,
			"platform" : "mercedesamgf1.com"
		}
	`)

	response, err := http.Post(server, "application/json", requestBody)
	trgiggerPanic(err)
	defer response.Body.Close()

	responseContent, err := io.ReadAll(response.Body)
	trgiggerPanic(err)

	var canvas strings.Builder
	_, err = canvas.Write(responseContent)
	trgiggerPanic(err)

	fmt.Println("response: \n", canvas.String())

}

func trgiggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
