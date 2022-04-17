package domain

type Look struct {
	Id               uint   `json:"-"`
	PhotoPath        string `json:"-"`
	OwnerId          uint   `json:"-"`
	Description      string `json:"description"`
	Season           string `json:"season"`
	TemperatureRange [2]int `json:"temperature_range"`
	Purpose          string `json:"purpose"`
	Priority         int    `json:"priority"`
}
