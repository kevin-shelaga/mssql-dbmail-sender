package main

import (
	"os"
	"strconv"

	"github.com/kevin-shelaga/mssql-dbmail-sender/helpers"
	"github.com/kevin-shelaga/mssql-dbmail-sender/logging"
	"github.com/kevin-shelaga/mssql-dbmail-sender/mail"

	"github.com/kevin-shelaga/mssql-dbmail-sender/sql"
)

func main() {

	logging.Information("***** Starting mssql-dbmail-sender *****")

	logging.Information("***** Getting appsettings.json *****")
	config := helpers.GetConfig("./appsettings.json")

	var s sql.DBMAIL = sql.T{ConnectionString: os.Getenv("GormConnectionString")}
	db := s.Connect()

	limit, _ := strconv.Atoi(helpers.Config["MailItemsLimit"].(string))
	sysmailMailItems := s.GetSysmailItems(db, limit)

	var m mail.SMTP = mail.T{Host: config["SmtpHost"].(string), Port: helpers.Config["SmtpPort"].(string), UserName: helpers.Config["SmtpUserName"].(string), Password: helpers.Config["SmtpPassword"].(string), From: helpers.Config["SmtpFrom"].(string)}

	for _, mailItem := range sysmailMailItems {
		logging.Information("**********")
		logging.Information("MailItemID: " + strconv.FormatUint(uint64(mailItem.MailItemID), 10))
		logging.Information("Recipients: " + mailItem.Recipients)
		logging.Information("CopyRecipients: " + mailItem.CopyRecipients)
		logging.Information("BlindCopyRecipients: " + mailItem.BlindCopyRecipients)
		logging.Information("Importance: " + mailItem.Importance)
		logging.Information("BodyFormat: " + mailItem.BodyFormat)
		logging.Information("Subject: " + mailItem.Subject)
		// logging.Information("Body: " + mailItem.Body)

		var errz error
		if helpers.Config["TestMode"].(bool) {
			err := m.Send(helpers.Config["TestModeRecipient"].(string), "", "", mailItem.Subject, mailItem.Body, mailItem.Importance, mailItem.BodyFormat)
			errz = err
		} else {
			if os.Getenv("ENVIRONMENT") == "Prodution" {
				err := m.Send(mailItem.Recipients, mailItem.CopyRecipients, mailItem.BlindCopyRecipients, mailItem.Subject, mailItem.Body, mailItem.Importance, mailItem.BodyFormat)
				errz = err
			}
		}

		s.UpdateSysmailItem(db, mailItem.MailItemID, errz)
	}

	logging.Information("***** Finished mssql-dbmail-sender *****")
}
