package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

type weatherData struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea      []NearestArea      `json:"nearest_area"`
}

type CurrentCondition struct {
	FeelsLike   string        `json:"FeelsLikeC"`
	Humidity    string        `json:"humidity"`
	Pressure    string        `json:"pressure"`
	Temperature string        `json:"temp_C"`
	WindSpeed   string        `json:"windspeedKmph"`
	WeatherDesc []WeatherDesc `json:"weatherDesc"`
	ObsTime     string        `json:"localObsDateTime"`
}

type WeatherDesc struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName   []AreaName `json:"areaName"`
	Country    []Country  `json:"country"`
	Latitude   string     `json:"latitude"`
	Longitude  string     `json:"longitude"`
	Population string     `json:"population"`
}

type AreaName struct {
	Value string `json:"value"`
}

type Country struct {
	Value string `json:"value"`
}

func main() {
	response := APIresponse(Input())
	instanceRfr := DataExtraction(response)
	fiilePtr := FileOperation(instanceRfr)
	InjectData(instanceRfr, fiilePtr)

	fmt.Println("Shell by Anshuman Pattnaik  Say hello: @github.com/ANSHPG")
	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Input() string {
	fmt.Print("Enter Location:- ")

	reader := bufio.NewReader(os.Stdin)
	ipContent, err := reader.ReadString('\n')
	triggerPanic(err)
	return strings.TrimSpace(ipContent)
}

func APIresponse(location string) []uint8 {
	server := fmt.Sprintf("https://wttr.in/%s?format=j1", location)
	fmt.Println("fetching data from server...")

	resolve, err := http.Get(server)
	triggerPanic(err)
	defer resolve.Body.Close()

	response, err := io.ReadAll(resolve.Body)
	triggerPanic(err)

	if !json.Valid(response) {
		log.Fatal("Invalid API response: type other that JSON")
	}

	return response
}

func DataExtraction(reponse []uint8) *weatherData {
	structReference := weatherData{}
	fmt.Println("Injecting data into new instance...")
	err := json.Unmarshal(reponse, &structReference)
	triggerPanic(err)

	return &structReference
}

func FileOperation(wdPtr *weatherData) *os.File {
	// fileStr := "./Ashft_weather_result/data.txt"
	location := (*wdPtr).NearestArea[0].AreaName[0].Value
	fileStr := fmt.Sprintf("./Ashft_weather_result/%s.txt", location)

	err := os.MkdirAll("./Ashft_weather_result", os.ModePerm)
	triggerPanic(err)

	file, err := os.Create(fileStr)
	triggerPanic(err)

	return file
}

func InjectData(wdPtr *weatherData, flPtr *os.File) {
	instance := *wdPtr

	var canvas strings.Builder

	location := instance.NearestArea[0].AreaName[0].Value
	myFigure := figure.NewFigure(location, "small", true)
	canvas.WriteString(myFigure.String())
	canvas.WriteString("<-----------ANSHift--Weather-Data----------->\n\n")

	fmt.Fprintf(&canvas, "Feels Like: %s°C\n", instance.CurrentCondition[0].FeelsLike)
	fmt.Fprintf(&canvas, "Humidity: %s\n", instance.CurrentCondition[0].Humidity)
	fmt.Fprintf(&canvas, "Pressure: %s\n", instance.CurrentCondition[0].Pressure)
	fmt.Fprintf(&canvas, "Temperature: %s°C\n", instance.CurrentCondition[0].Temperature)
	fmt.Fprintf(&canvas, "Wind Speed: %sKm/Hr\n", instance.CurrentCondition[0].WindSpeed)
	fmt.Fprintf(&canvas, "Weather Description: %s\n", instance.CurrentCondition[0].WeatherDesc[0].Value)
	fmt.Fprintf(&canvas, "Observation Time: %s\n", instance.CurrentCondition[0].ObsTime)

	fmt.Fprintf(&canvas, "Location: %s\n", location)
	fmt.Fprintf(&canvas, "Country: %s\n", instance.NearestArea[0].Country[0].Value)

	fmt.Fprintf(&canvas, "Latitude: %s\n", instance.NearestArea[0].Latitude)
	fmt.Fprintf(&canvas, "Longitude: %s\n", instance.NearestArea[0].Longitude)
	fmt.Fprintf(&canvas, "Population: %s\n", instance.NearestArea[0].Population)

	_, err := io.WriteString(flPtr, canvas.String())
	triggerPanic(err)
	defer flPtr.Close()

	fmt.Println(canvas.String())
}
