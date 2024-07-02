package model

type Car struct {
	ID           int64  `json:"id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Mileage      int64  `json:"mileage"`
	OwnersNumber int    `json:"onwers_number"`
}
