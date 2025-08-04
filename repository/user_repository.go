package repository

import (
	"database/sql"
	"log"
	"notification-api/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var password string

	passwordhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to create password", err)
		return 0, nil
	}

	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" + "(email, password_hash)" + "VALUES ($1, $2) RETURNING id_user")
	if err != nil {
		log.Println("Failed to create new user:", err)
		return 0, err
	}
	err = query.QueryRow(user.Email, string(passwordhash)).Scan(&id)
	if err != nil {
		log.Println("Failed to allocate new user:", err)
		return 0, err
	}
	query.Close()

	return id, nil
}
