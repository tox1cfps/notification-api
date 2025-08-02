package model

type Notification struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	EmailAuthor   string `json:"emailauthor"`
	EmailReceiver string `json:"emailreceiver"`
}
