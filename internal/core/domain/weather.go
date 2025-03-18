package domain

type Weather struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Description string `json:"description"`
	Humidity    string `json:"humidity"`
	WindSpeed   string `json:"wind_speed"`
}
