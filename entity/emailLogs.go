package entity

import "time"

type EmailLog struct {
	ID        int
	ToEmail   string
	Subject   string
	Body      string
	Status    string
	Delivered bool
	ErrorText *string
	CreatedAt time.Time
}

type EmailLogResponse struct {
	ID        int     `json:"id"`
	ToEmail   string  `json:"to_email"`
	Subject   string  `json:"subject"`
	Body      string  `json:"body"`
	Status    string  `json:"status"`
	Delivered bool    `json:"delivered"`
	ErrorText *string `json:"error_text"`
	CreatedAt string  `json:"created_at"`
}
