package model

import "time"

//Response represent the basic response object
type Response struct {
	Status     string      `json:"status" example:"OK"`
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"success fetch data"`
	Timestamp  time.Time   `json:"timestamp" example:"1970-01-01T00:00:00.000000000Z"`
	Data       interface{} `json:"data"`
}
