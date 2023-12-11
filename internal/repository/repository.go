package repository

import "gorm.io/gorm"

func NewTicketRepository(mysql *gorm.DB) TicketRepository {
	return &ticketRepository{mysql: mysql}
}
