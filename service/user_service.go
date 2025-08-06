package service

import (
	"log"
	"notification-api/model"
	"notification-api/repository"
)

type UserService struct {
	repository          repository.UserRepository
	NotificationService NotificationService
}

func NewUserService(repo repository.UserRepository, notificationService NotificationService) UserService {
	return UserService{
		repository:          repo,
		NotificationService: notificationService,
	}
}
func (us *UserService) CreateUser(user model.User) (model.User, error) {
	userID, err := us.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = userID

	notification := model.Notification{
		Title:         "Bem vindo ao nosso site!",
		Content:       "Olá, sua conta foi criada com sucesso!",
		EmailAuthor:   "arthurrodriguesfonseca@gmail.com",
		EmailReceiver: user.Email,
	}

	_, err = us.NotificationService.CreateNotification(notification)
	if err != nil {
		log.Println("Failed to create notification:", err)
	}

	return user, nil
}
