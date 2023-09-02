package model

type Attraction struct {
	ID         uint    `gorm:"PrimaryKey" json:"id"`
	Name       string  `json:"name"`
	Coverimage string  `json:"coverimage"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
}
