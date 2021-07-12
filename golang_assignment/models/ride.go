package models

type Ride struct {
	r_id           int     `json:"r_id" grom:"PRIMARY_KEY"`
	dr_id          int     `json:"dr_id"`
	starting_point string  `json:"starting_point"`
	ending_point   string  `json:"ending_point"`
	distance       float32 `json:"distance"`
	travel_time    string  `json:"travel_time"`
	pickup_time    string  `json:"pickup_time"`
	createdAt      string  `json:"createdAt"`
	updatedAt      string  `json:"updatedAt"`
}
