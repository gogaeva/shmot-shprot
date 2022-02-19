package model

type Look struct {
	Id               int    `json:"-"`
	PhotoPath        string `json:"-"`
	OwnerId          int    `json:"-"`
	Description      string `json:"description"`
	Season           string `json:"season"`
	TemperatureRange [2]int `json:"temperature_range"`
	Purpose          string `json:"purpose"`
	Priority         int    `json:"priority"`
}
