package repository

import (
	"database/sql"
	"log"
	"notification-api/model"
)

type NotificationRepository struct {
	connection *sql.DB
}

func NewNotificationRepository(connection *sql.DB) NotificationRepository {
	return NotificationRepository{
		connection: connection,
	}
}

func (nr *NotificationRepository) CreateNotification(notification model.Notification) (int, error) {
	var id int
	query, err := nr.connection.Prepare("INSERT INTO notifications" + "(title, content, emailauthor, emailreceiver)" + "VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = query.QueryRow(notification.Title, notification.Content, notification.EmailAuthor, notification.EmailReceiver).Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
