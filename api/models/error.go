package models

type JSONError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
