package service

import (
	"crypto/rand"
	"errors"
	"io"
	"log"
	"notification-api/model"
	"notification-api/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordService struct {
	ResetRepo           repository.ResetPasswordRepository
	UserRepo            repository.UserRepository
	NotificationService NotificationService
}

func NewResetPasswordService(resetRepo repository.ResetPasswordRepository, userRepo repository.UserRepository, notificationService NotificationService) ResetPasswordService {
	return ResetPasswordService{
		ResetRepo:           resetRepo,
		UserRepo:            userRepo,
		NotificationService: notificationService,
	}
}

func (rps *ResetPasswordService) ResetPassword(email string) {
	user, err := rps.UserRepo.GetUserByEmail(email)
	if err != nil {
		log.Println("Failed to find user", err)
		return
	}

	if user.ID == 0 {
		log.Println("User not found")
		return
	}
	// criar token
	token := EncodeToString(6)

	reset := model.PasswordReset{
		Email:     email,
		Token:     token,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	err = rps.ResetRepo.SaveResetToken(reset)
	if err != nil {
		log.Println("Failed to save token:", err)
		return
	}

	notification := model.Notification{
		Title:         "Recuperação de senha",
		Content:       "Seu código de recuperação é: " + token,
		EmailAuthor:   "arthurrodriguesfonseca@gmail.com",
		EmailReceiver: email,
	}

	_, err = rps.NotificationService.CreateNotification(notification)
	if err != nil {
		log.Println("Failed to send notification:", err)
	}
}

func (rps *ResetPasswordService) ValidUser(email, token, newPassword string) error {
	reset, err := rps.ResetRepo.GetToken(email, token)
	if err != nil {
		return err
	}

	if reset == nil {
		return errors.New("invalid token")
	}

	if time.Now().UTC().After(reset.ExpiresAt.UTC()) {
		return errors.New("token expired")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = rps.UserRepo.UpdatePasswordByEmail(email, string(hashed))
	if err != nil {
		return err
	}

	err = rps.ResetRepo.DeleteToken(email)
	if err != nil {
		log.Println("Failed to delete token, but password was updated")
	}

	notification := model.Notification{
		Title:         "Senha alterada com sucesso",
		Content:       "Heyey, sua senha foi alterada recentemente. Se não foi você, entre em contato conosco imediatamente",
		EmailAuthor:   "arthurrodriguesfonseca@gmail.com",
		EmailReceiver: email,
	}

	_, err = rps.NotificationService.CreateNotification(notification)
	if err != nil {
		log.Println("Failed to send confirmation email:", err)
	}

	return nil
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
