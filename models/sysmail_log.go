package models

import (
	"time"
)

//SysmailLog sysmail_log table
type SysmailLog struct {
	LogID       uint      `gorm:"column:log_id;primaryKey"`
	EventType   uint      `gorm:"column:event_type"`
	LogDate     time.Time `gorm:"column:log_date"`
	Description string    `gorm:"column:description"`
	ProcessID   uint      `gorm:"column:process_id"`
	MailitemID  uint      `gorm:"column:mailitem_id"`
	LastModDate time.Time `gorm:"column:last_mod_date"`
	LastModUser string    `gorm:"column:last_mod_user"`
}

//TableName set the table name
func (SysmailLog) TableName() string {
	return "sysmail_log"
}
