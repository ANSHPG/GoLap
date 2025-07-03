package main

import "fmt"

func main() {
	fmt.Println("Function like a Machine, Roar like a Beast!")
	log()
	fmt.Println(lap2(75, 90))
	_, msg := proDriver(2, 5, 8, 9)
	fmt.Println(msg)
	// not allowed to write functions inside main() or used defined Function
}

func log() {
	fmt.Println("Hamilton sweeped pass Max to P1!")
}

func lap2(time1, time2 int) int {
	return time1 + time2
}

func proDriver(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "BOX"
}
