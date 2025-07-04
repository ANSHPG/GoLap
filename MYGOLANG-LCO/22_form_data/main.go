package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("sharing track records...")
	FormData()
}

func FormData() {
	const server string = "https://httpbin.org/post"

	detail := url.Values{}
	detail.Add("Company", "AMG")
	detail.Add("points", "203")
	detail.Add("position", "3rd")

	resolve, err := http.PostForm(server, detail)
	triggerPanic(err)
	defer resolve.Body.Close()

	response, err := io.ReadAll(resolve.Body)
	triggerPanic(err)

	var canvas strings.Builder
	_, err = canvas.Write(response)
	triggerPanic(err)

	fmt.Println("response: \n", canvas.String())

}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
