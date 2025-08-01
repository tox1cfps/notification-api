package repository

import "database/sql"

type NotificationRepository struct {
	connection *sql.DB
}

func NewNotificationRepository(connection *sql.DB) NotificationRepository {
	return NotificationRepository{
		connection: connection,
	}
}
