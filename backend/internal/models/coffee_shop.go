package models

type Cafe struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Rating           float64 `json:"rating"`
	HoursOfOperation string  `json:"hoursOfOperation"`
	HasWifi          bool    `json:"hasWifi"`
	HasOutlets       bool    `json:"hasOutlets"`
	IsIndependent    bool    `json:"isIndependent"`
	PhotoURL         string  `json:"photoUrl"`
}
