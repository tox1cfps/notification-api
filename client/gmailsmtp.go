package client

import (
	"fmt"
	"net/smtp"
	"notification-api/config"
)

type GmailsmtpClient struct{}

func NewGmailsmtpClient() GmailsmtpClient {
	return GmailsmtpClient{}
}

func (g GmailsmtpClient) SendEmail(toEmail string, subject string, body string) error {
	settings := config.Settings
	from := settings.Gmailsmtp.Email
	password := settings.Gmailsmtp.Password

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	to := []string{toEmail}

	message := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
			"MIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
		from, toEmail, subject, body,
	))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Email enviado com sucesso!")
	return nil
}
