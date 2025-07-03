package main

import (
	"fmt"
	"net/url"
)

// const MSurl string = "https://f1api.dev/api/drivers/search?q=lewis"
const MSurl string = "https://api.openf1.org/v1/drivers?driver_number=44&session_key=9158"

func main() {
	fmt.Println("United Relay Laps")
	fmt.Println(MSurl)

	result, _ := url.Parse(MSurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("query collection: %+v\n", qparams)
	fmt.Println(qparams["driver_number"], "---", qparams["session_key"])

	parseOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "api.openf1.org",
		Path:    "/v1/drivers",
		RawPath: "driver_number=44&session_key=9158",
	}

	nextUrl := parseOfUrl.String()
	fmt.Println(nextUrl)
}
