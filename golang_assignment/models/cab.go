package models

type Cab struct {
	cab_id      int     `json:"cab_id"`
	driver_name string  `json:"driver_name"`
	model       string  `json:"model"`
	reg_no      string  `json:"reg_no"`
	status      string  `json:"status"`
	latitude    float32 `json:"latitude"`
	longitude   float32 `json:"longitude"`
	createdAt   string  `json:"createdAt"`
	updatedAt   string  `json:"updatedAt"`
}
