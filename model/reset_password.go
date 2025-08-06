package model

import "time"

type PasswordReset struct {
	Token     string
	Email     string
	ExpiresAt time.Time `json:"expires_at"`
}
