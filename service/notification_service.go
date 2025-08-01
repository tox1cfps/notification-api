package service

import (
	"notification-api/client"
	"notification-api/repository"
)

type NotificationService struct {
	NotificationRepository repository.NotificationRepository
	Gmailsmtp              client.GmailsmtpClient
}

func NewNotificationService(repo repository.NotificationRepository, g client.GmailsmtpClient) NotificationService {
	return NotificationService{
		NotificationRepository: repo,
		Gmailsmtp:              g,
	}
}
func (n NotificationService) SendEmailToUser() {

}
