package model

type User struct {
	ID            int    `json:"id_user"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Password_Hash string `json:"-"`
}
