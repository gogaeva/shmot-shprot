package model

type Cloth struct {
	Id    int    `json:"-"`
	Class string `json:"class"`
	Brand string `json:"brand"`
	Color int    `json:"color"`
}
