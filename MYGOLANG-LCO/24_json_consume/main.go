package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type weatherData struct {
	CurrentCondition []CurrentCondion `json:"current_condition"`
}

type CurrentCondion struct {
	TempC   string `json:"temp_C"`
	TempF   string `json:"temp_F"`
	ObsTime string `json:"localObsDateTime"`
}

func main() {
	fmt.Println("recieving data from AMG server:1 ...")
	ApiManagement()
}

func DecodeJson(response []uint8) {
	// jsonResponse := []byte()
	// var responseJson map[string]interface{}
	structReference := weatherData{}

	err := json.Unmarshal(response, &structReference)
	triggerPanic(err)

	fmt.Printf("%+v\n", structReference)
	floatC, err := strconv.ParseFloat(structReference.CurrentCondition[0].TempC, 32)
	triggerPanic(err)
	fmt.Printf("%f", floatC)
}

func ApiManagement() {
	const server string = "https://wttr.in/monaco?format=j1"
	// fmt.Sprintf() => for string injection within varible

	resolve, err := http.Get(server)
	triggerPanic(err)
	defer resolve.Body.Close()

	response, err := io.ReadAll(resolve.Body)
	triggerPanic(err)

	// fmt.Printf("%T", response)
	DecodeJson(response)
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
