package main

import (
	"fmt"
	"sort"
)

func remove(arr *[]string, index *int) {
	*arr = append((*arr)[:*index], (*arr)[*index+1:]...)
}

func main() {
	fmt.Println("Slice the track, before your tires cool down!")

	var franchise = []string{"McLaren", "AstonMartin", "LaFerrari", "MercedesAMG"}
	fmt.Printf("type:- %T\n", franchise)

	franchise = append(franchise, "RedBull", "Haas")
	fmt.Println(franchise)

	// franchise = append(franchise[0:3])
	fmt.Println(franchise)

	pts := make([]int, 4)

	// actual team standings with points acc. to 07-01-25 20:10 +5:30 GMT
	pts[0] = 417
	pts[1] = 28
	pts[2] = 210
	pts[3] = 209
	pts = append(pts, 162, 29) // avoid crash

	fmt.Println(pts)
	sort.Ints(pts)
	fmt.Println(pts)
	fmt.Println(sort.IntsAreSorted(pts))

	// how to remove a value from slices based on index
	var super = []string{"AMG", "Audi", "Porsche", "Ferrari", "Lambourghini"}
	fmt.Println(super)

	var index int = 2
	super = append(super[:index], super[index+1:]...)
	fmt.Println(super)

	remove(&super, &index)
	fmt.Println(super)
}
