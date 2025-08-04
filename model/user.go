package model

type User struct {
	ID            int    `json:"id_user"`
	Email         string `json:"email"`
	Password_Hash string `json:"password_hash"`
}
