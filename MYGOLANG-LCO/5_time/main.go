package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Reduce time for the lap")
	prsntTime := time.Now()
	fmt.Println(prsntTime)

	fmt.Println(prsntTime.Format("01-02-2006"))
	fmt.Println(prsntTime.Format("15:04:05"))
	fmt.Println(prsntTime.Format("Monday"))

	createDate := time.Date(2017, time.October, 18, 3, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006"))
	fmt.Println(strings.Split(createDate.Format("15:04:05"), ":")[0])
	fmt.Println(createDate.Format("Monday"))

	fmt.Scanln()
}

//  to bulid it
// - go env [ to know about documents]
// - GOOS="windows" go build , its diff. for diff. os

// also check out - pkg.go.dev/runtime
