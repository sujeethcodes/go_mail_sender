package service

import (
	"fmt"
	"go_mail_sender/entity"
	"net/smtp"
	"os"
)

type MailService struct{}

func (s *MailService) Send(emailData entity.Mail) (*entity.EmailLog, error) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	msg := []byte(fmt.Sprintf("Subject: %s\r\nTo: %s\r\n\r\n%s", emailData.Subject, emailData.To, emailData.Body))

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, smtpUser, []string{emailData.To}, msg)

	// build log
	log := &entity.EmailLog{
		ToEmail:   emailData.To,
		Subject:   emailData.Subject,
		Body:      emailData.Body,
		Status:    "sent_attempted",
		Delivered: err == nil,
	}
	if err != nil {
		errMsg := err.Error()
		log.ErrorText = &errMsg
	}

	return log, err
}
