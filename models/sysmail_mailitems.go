package models

import (
	"time"
)

//SysmailMailItems sysmail_log table
type SysmailMailItems struct {
	MailItemID              uint      `gorm:"column:mailitem_id;primaryKey"`
	ProfileID               uint      `gorm:"column:profile_id"`
	Recipients              string    `gorm:"column:recipients"`
	CopyRecipients          string    `gorm:"column:copy_recipients"`
	BlindCopyRecipients     string    `gorm:"column:blind_copy_recipients"`
	Subject                 string    `gorm:"column:subject"`
	FromAddress             string    `gorm:"column:from_address"`
	ReplyTo                 string    `gorm:"column:reply_to"`
	Body                    string    `gorm:"column:body"`
	BodyFormat              string    `gorm:"column:body_format"`
	Importance              string    `gorm:"column:importance"`
	Sensitivity             string    `gorm:"column:sensitivity"`
	FileAttachments         string    `gorm:"column:file_attachments"`
	AttachmentEncoding      string    `gorm:"column:attachment_encoding"`
	Query                   string    `gorm:"column:query"`
	ExecuteQueryDatabase    string    `gorm:"column:execute_query_database"`
	AttachQueryResultAsFile bool      `gorm:"column:attach_query_result_as_file"`
	QueryResultHeader       bool      `gorm:"column:query_result_header"`
	QueryResultWidth        uint      `gorm:"column:query_result_width"`
	QueryResultSeparator    string    `gorm:"column:query_result_separator"`
	ExcludeQueryOutput      bool      `gorm:"column:exclude_query_output"`
	AppendQueryError        bool      `gorm:"column:append_query_error"`
	SendRequestDate         time.Time `gorm:"column:send_request_date"`
	SendRequestUser         string    `gorm:"column:send_request_user"`
	SentAccountID           uint      `gorm:"column:sent_account_id"`
	SentStatus              uint      `gorm:"column:sent_status"`
	SentDate                time.Time `gorm:"column:sent_date"`
	LastModDate             time.Time `gorm:"column:last_mod_date"`
	LastModUser             string    `gorm:"column:last_mod_user"`
}

//TableName set the table name
func (SysmailMailItems) TableName() string {
	return "sysmail_mailitems"
}
