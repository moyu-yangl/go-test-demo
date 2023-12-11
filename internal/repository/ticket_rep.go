package repository

import (
	"go-test-demo/constants"
	"go-test-demo/internal/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetTicketById(id uint64) (models.Ticket, error)
	PageNormalTicketRecord(condition map[string]interface{}, offset int, limit int, orders string) ([]models.Ticket, error)
	ListTicketByProjectId(projectId int64) ([]models.Ticket, error)
}

type ticketRepository struct {
	mysql *gorm.DB
}

func (rep *ticketRepository) GetTicketById(id uint64) (models.Ticket, error) {
	records := make([]models.Ticket, 0)
	tx := rep.mysql.Where(&models.Ticket{Id: id}).Find(&records)
	if len(records) > 0 {
		return records[0], nil
	} else {
		return constants.NilTicket, tx.Error
	}
}
func (rep *ticketRepository) PageNormalTicketRecord(condition map[string]interface{}, offset int, limit int, orders string) ([]models.Ticket, error) {
	records := make([]models.Ticket, 0)
	tx := rep.mysql.Where(condition).Offset(offset).Limit(limit).Order(orders).Find(&records)
	return records, tx.Error
}
func (rep *ticketRepository) ListTicketByProjectId(projectId int64) ([]models.Ticket, error) {
	records := make([]models.Ticket, 0)
	tx := rep.mysql.Where(models.Ticket{ProjectId: projectId}).Find(&records)

	return records, tx.Error
}
