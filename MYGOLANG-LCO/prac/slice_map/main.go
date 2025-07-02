package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Review the Map, Slice the Track!")

	var linear = []string{"McLaren", "Ferrari", "Mercedes", "RedBullRacing", "AstonMartin"}
	linear = append(linear, "Haas")
	// fmt.Println(linear)
	linear = remove(linear, slices.Index(linear, "AstonMartin"))
	fmt.Println(linear)

	franchise := make(map[string]int32)
	integrate(&franchise)
	fmt.Println(franchise)

}

func remove(arr []string, index int) []string {
	return append(arr[:index], arr[index+1:]...)
}

func integrate(track *map[string]int32) {
	(*track)["McLaren"] = 417
	(*track)["Ferrari"] = 210
	(*track)["Mercedes"] = 209
	(*track)["RedBullRacing"] = 162
	(*track)["Haas"] = 29
	(*track)["AstonMartin"] = 28
}
