package main

import "fmt"

func main() {
	fmt.Println("Review Map of the Track!")

	franchise := make(map[string]uint32)

	// F1 team points, 07-02-25 11:50 +5:30 GMT
	franchise["McLaren"] = 417
	franchise["Ferrari"] = 210
	franchise["Mercedes"] = 209
	franchise["RedBullRacing"] = 162
	franchise["Haas"] = 29
	franchise["AstonMartin"] = 28

	// fmt.Println(franchise)
	key := "McLaren"
	// fmt.Printf("key: %s\nvalue: %d\n", key, franchise[key])

	delete(franchise, key)
	// fmt.Println(franchise)

	// Loops in GO

	for key, val := range franchise {
		fmt.Printf("key: %s,  value: %d\n", key, val)
	}

	for _, val := range franchise {
		fmt.Printf("value: %d\n", val)
	}

}
