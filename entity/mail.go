package entity

type Mail struct {
	To      string `gorm:"size:255"`
	Subject string `gorm:"size:512"`
	Body    string `gorm:"type:text"`
}

type MailSendResponse struct {
	Status    string  `json:"status"`
	Delivered bool    `json:"delivered"`
	Error     *string `json:"error"`
}
