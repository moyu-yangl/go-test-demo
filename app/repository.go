package app

import (
	"go-test-demo/internal/repository"
	"gorm.io/gorm"
	"sync"
)

var ticketRepository repository.TicketRepository

var onceRep sync.Once

func InitRepository(mysql *gorm.DB) error {
	onceRep.Do(func() {
		ticketRepository = repository.NewTicketRepository(mysql)
	})
	return nil
}

func TicketRepository() repository.TicketRepository {
	return ticketRepository
}
