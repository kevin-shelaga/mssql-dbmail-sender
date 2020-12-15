package mail

import (
	"net/smtp"

	"github.com/kevin-shelaga/mssql-dbmail-sender/logging"

	"github.com/kevin-shelaga/mssql-dbmail-sender/helpers"
)

//SMTP interace for smtp package
type SMTP interface {
	Send(to string, cc string, bcc string, subject string, body string, importance string, format string) error
}

//T is the interface struct type
type T struct {
	Host     string
	Port     string
	UserName string
	Password string
	From     string
}

const (
	lowImportance    = "0"
	normalImportance = "1"
	highImportance   = "2"
	textFormat       = "text/plain; charset=\"UTF-8\""
	htmlFormat       = "text/html; charset=\"UTF-8\""
)

//Send smtp email
func (t T) Send(to string, cc string, bcc string, subject string, body string, importance string, format string) error {

	logging.Information("Sending mail...")

	var err error
	var server = t.Host + ":" + t.Port

	auth := smtp.PlainAuth("", t.UserName, t.Password, t.Host)

	var imprt = ""
	if importance == "HIGH" {
		imprt = highImportance
	} else if importance == "NORMAL" {
		imprt = normalImportance
	} else if importance == "LOW" {
		imprt = lowImportance
	}

	var frmt = ""
	if format == "TEXT" {
		frmt = textFormat
	} else if format == "HTML" {
		frmt = htmlFormat
	}

	msg := helpers.ComposeMimeMail(to, t.From, subject, body, imprt, frmt)

	err = smtp.SendMail(server, auth, t.From, []string{to}, msg)
	if err != nil {
		logging.Critical(err)
	} else {
		logging.Information("Mail sent!")
	}

	return err
}
