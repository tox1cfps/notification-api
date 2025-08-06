package repository

import (
	"database/sql"
	"notification-api/model"
)

type ResetPasswordRepository struct {
	connection *sql.DB
}

func NewResetPasswordRepository(connection *sql.DB) ResetPasswordRepository {
	return ResetPasswordRepository{
		connection: connection,
	}
}

func (rpr *ResetPasswordRepository) SaveResetToken(reset model.PasswordReset) error {
	query := `
		INSERT INTO password_resets (email, token, expires_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO UPDATE
		SET token = EXCLUDED.token,
		    expires_at = EXCLUDED.expires_at;
	`
	_, err := rpr.connection.Exec(query, reset.Email, reset.Token, reset.ExpiresAt)
	return err
}

func (rpr *ResetPasswordRepository) GetToken(email, token string) (*model.PasswordReset, error) {
	var reset model.PasswordReset
	query := "SELECT email, token, expires_at FROM password_resets WHERE email = $1 AND token = $2"
	err := rpr.connection.QueryRow(query, email, token).Scan(&reset.Email, &reset.Token, &reset.ExpiresAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &reset, nil
}

func (rpr *ResetPasswordRepository) DeleteToken(email string) error {
	_, err := rpr.connection.Exec("DELETE FROM password_resets WHERE email = $1", email)
	return err
}
