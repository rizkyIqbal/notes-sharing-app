package models

import "time"

type Log struct {
	ID           int       `json:"id"`
	Datetime     time.Time `json:"datetime"`
	Method       string    `json:"method"`
	Endpoint     string    `json:"endpoint"`
	Headers      string    `json:"headers"`
	Payload      string    `json:"payload"`
	ResponseBody string    `json:"response_body"`
	StatusCode   int       `json:"status_code"`
}
