package repository

import (
	"database/sql"
	"fmt"
	"log"
	"notification-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	query := "SELECT id_user, email, password_hash FROM users WHERE email = $1"
	row := ur.connection.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email, &user.Password_Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found with this email:", email)

			return model.User{}, fmt.Errorf("user not found")
		}
		log.Println("Failed to get user by email:", err)
		return model.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query, err := ur.connection.Prepare(`
        INSERT INTO users (email, password_hash)
        VALUES ($1, $2) RETURNING id_user
    `)
	if err != nil {
		log.Println("Failed to create new user:", err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(user.Email, user.Password_Hash).Scan(&id)
	if err != nil {
		log.Println("Failed to allocate new user:", err)
		return 0, err
	}
	return id, nil
}

func (ur *UserRepository) UpdatePasswordByEmail(email, newPasswordHash string) error {
	query := "UPDATE users SET password_hash = $1 WHERE email = $2"
	_, err := ur.connection.Exec(query, newPasswordHash, email)
	return err
}
