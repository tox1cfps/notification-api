package service

import (
	"notification-api/model"
	"notification-api/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repository: repo,
	}
}
func (us *UserService) CreateUser(user model.User) (model.User, error) {
	userID, err := us.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = userID

	return user, nil
}
