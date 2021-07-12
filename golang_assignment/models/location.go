package models

type Location struct {
	loc_id    int     `json:"loc_id" gorm:"PRIMARY_KEY"`
	name      string  `json:"name"`
	latitude  float32 `json:"latitude"`
	longitude float32 `json:"longitude"`
	createdAt string  `json:"createdAt"`
	updatedAt string  `json:"updatedAt"`
}
