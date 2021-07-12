package models

type Payment struct {
	ride_id   int     `json:"ride_id"`
	p_type    string  `json:"type"`
	cost      float32 `json:"cost"`
	status    string  `json:"status"`
	createdAt string  `json:"createdAt"`
	updatedAt string  `json:"updatedAt"`
}
