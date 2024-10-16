package models

type Cafe struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Address          string           `json:"address"`
	Latitude         float64          `json:"latitude"`
	Longitude        float64          `json:"longitude"`
	Rating           float64          `json:"rating"`
	HoursOfOperation WeekOpeningHours `json:"hoursOfOperation"`
	HoursLastUpdated string           `json:"hoursLastUpdated"`
	HasWifi          bool             `json:"hasWifi"`
	HasOutlets       bool             `json:"hasOutlets"`
	IsIndependent    bool             `json:"isIndependent"`
	PhotoURL         string           `json:"photoUrl"`
}

type DayHours struct {
	Day       string `json:"day"`
	OpenTime  string `json:"openTime"`
	CloseTime string `json:"closeTime"`
	IsClosed  bool   `json:"isClosed"`
	Is24Hours bool   `json:"is24Hours"`
}

type WeekOpeningHours struct {
	Monday    DayHours `json:"monday"`
	Tuesday   DayHours `json:"tuesday"`
	Wednesday DayHours `json:"wednesday"`
	Thursday  DayHours `json:"thursday"`
	Friday    DayHours `json:"friday"`
	Saturday  DayHours `json:"saturday"`
	Sunday    DayHours `json:"sunday"`
}
