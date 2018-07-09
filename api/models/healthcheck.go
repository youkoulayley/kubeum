package models

type Healthcheck struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
