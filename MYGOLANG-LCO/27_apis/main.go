package main

type Company struct {
	CompanyId uint8    `json:"c_id"`
	FullName  string   `json:"full_name"`
	Base      string   `json:"base"`
	TeamChief string   `json:"team_chief"`
	PowerUnit string   `json:"power_unit"`
	Driver1   []Driver `json:"driver_1"`
	Driver2   []Driver `json:"driver_2"`
}

type Driver struct {
	DiverId   uint8  `json:"d_id"`
	GrandPrix uint16 `json:"grand_prix"`
	Podiums   uint16 `json:"podiums"`
	Wins      uint16 `json:"wins"`
}

func (c *Company) IsEmpty() bool {
	return c.CompanyId == 0 && c.FullName == "     "
}

func main() {

}
