package service

import (
	"notification-api/client"
	"notification-api/model"
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
func (ns NotificationService) CreateNotification(notification model.Notification) (model.Notification, error) {
	notificationID, err := ns.NotificationRepository.CreateNotification(notification)
	if err != nil {
		return model.Notification{}, err
	}

	notification.ID = notificationID

	return notification, nil
}
