package main

import (
	"encoding/json"
	"fmt"
)

// there should be no space in json allias
type teamProfile struct {
	TeamName string  `json:"team_name"`
	NetWorth float32 `json:"-"`
	// "-" to hide information json parsing usally for passwords
	TeamChief string   `json:"team_chief"`
	PowerUnit string   `json:"power_unit"`
	Drivers   []string `json:"drivers,omitempty"`
}

func main() {
	fmt.Println("Json Pars-es Hamilton")
	EncodeJson()
}

func EncodeJson() {
	//  data is taken from officail F1 site => 07-04-25 8:39 +5:30 GMT
	tPSlice := []teamProfile{
		{
			"Mercedes-AMG PETRONAS Formula One Team",
			3.8,
			"Toto Wolff",
			"Mercedes",
			[]string{"George Russell", "Kimi Antonelli"},
		},

		{
			"McLaren Formula 1 Team",
			2.2,
			"Andrea Stella",
			"Mercedes",
			[]string{"Oscar Piastri", "Lando Norris"},
		},

		{
			"Oracle Red Bull Racing",
			2.6,
			"Christian Horner",
			"Honda RBPT",
			[]string{"Max Verstappen", "Yuki Tsunoda"},
		},

		{
			"Scuderia Ferrari HP",
			3.9,
			"Frédéric Vasseur",
			"Ferrari",
			[]string{"Charles Leclerc", "Lewis Hamilton"},
		},

		{
			"Aston Martin Aramco Formula One Team",
			1.375,
			"Andy Cowell",
			"Mercedes",
			[]string{"Fernando Alonso", "Felipe Drugovich"},
		},
	}

	jsonPKG, err := json.MarshalIndent(tPSlice, "", "\t")
	triggerPanic(err)

	fmt.Printf("%s\n", jsonPKG)
}

func triggerPanic(err error) {
	if err != nil {
		panic(err)
	}
}
