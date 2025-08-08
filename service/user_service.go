package service

import (
	"fmt"
	"log"
	"notification-api/model"
	"notification-api/repository"

	"golang.org/x/crypto/bcrypt"
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

func (us *UserService) LoginUser(email, password string) (model.User, error) {
	user, err := us.repository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("wrong Password")
	}

	return user, nil
}
