package domain

type Cloth struct {
	Id      uint   `json:"-"`
	PhotoId string `json:"-"`
	OwnerId uint   `json:"owner_id"`
	Class   string `json:"class"`
	Brand   string `json:"brand"`
	Color   int    `json:"color"`
}
