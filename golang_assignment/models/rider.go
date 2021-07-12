package models

type Rider struct {
	cust_id   int    `json:"cust_id" grom:"PRIMARY_KEY"`
	password  string `json:"password"`
	name      string `json:"name"`
	phone     int    `json:"phone"`
	createdAt string `json:"createdAt"`
	updatedAt string `json:"updatedAt"`
}
