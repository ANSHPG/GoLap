package main

import "fmt"

func main() {
	fmt.Println("Structre your Statergy to the Podium!")
	// no inheritance, super ot parent in goLang

	driver := Career{"Lewis_Hamilton", 44, 367, 202, 7}
	fmt.Printf("Career Stats:-\n%+v\n", driver)
	fmt.Printf("Name:- %v\nDriver Number:- %d", driver.Name, driver.Number)
}

type Career struct {
	Name          string
	Number        int32
	GrandPrix     int64
	Podiums       int32
	WorldChampion uint16
}
