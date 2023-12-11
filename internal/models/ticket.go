package models

type Ticket struct {
	Id          uint64 `json:"id" gorm:"primaryKey"`
	TicketNo    string `json:"ticketNo"`
	Type        string `json:"type"`
	ProjectId   int64  `json:"projectId"`
	BusinessNo  string `json:"businessNo"`
	UserId      int64  `json:"userId"`
	AuditResult string `json:"auditResult"`
	Remark      string `json:"remark"`
}

func (_ Ticket) TableName() string {
	return "ticket"
}
