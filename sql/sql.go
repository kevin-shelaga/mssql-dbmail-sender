package sql

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/kevin-shelaga/mssql-dbmail-sender/logging"
	"github.com/kevin-shelaga/mssql-dbmail-sender/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//DBMAIL interace for accessing dbmail tables in NewsDesk
type DBMAIL interface {
	Connect() *gorm.DB
	GetSysmailItems(db *gorm.DB, limit int) []models.SysmailMailItems
	UpdateSysmailItem(db *gorm.DB, mailItemID uint, err error)
	GetSysmailItemsCount(db *gorm.DB) int64
}

//T is the interface struct type
type T struct {
	ConnectionString string
}

const (
	unsent   = 0
	sent     = 1
	failed   = 2
	retrying = 3
)

//Connect returns a new gorm db
func (t T) Connect() *gorm.DB {

	logging.Information("Connecting to database...")

	dsn := t.ConnectionString
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true}})

	if err != nil {
		logging.Critical(err)
		os.Exit(1)
	}

	logging.Information("Connected to database!")

	return db
}

//GetSysmailItems returns a list of sysmailMailitems
func (t T) GetSysmailItems(db *gorm.DB, limit int) []models.SysmailMailItems {

	logging.Information("Getting sysmail_mailitems...")

	var sysmailMailitems []models.SysmailMailItems

	result := db.Limit(limit).Where(map[string]interface{}{"sent_status": unsent}).Find(&sysmailMailitems)

	logging.Information("Got " + strconv.FormatInt(result.RowsAffected, 10) + " sysmail_mailitems!")

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logging.Critical(result.Error)
		os.Exit(1)
	}

	return sysmailMailitems
}

//UpdateSysmailItem updates sysmailMailitems and sysmailLog
func (t T) UpdateSysmailItem(db *gorm.DB, mailItemID uint, err error) {
	logging.Information("Update sysmail_mailitems...")

	var sysmailMailitems models.SysmailMailItems
	var sysmailLog models.SysmailLog

	if err != nil {
		db.Model(&sysmailMailitems).Where(models.SysmailMailItems{MailItemID: mailItemID}).Updates(models.SysmailMailItems{SentStatus: uint(failed), LastModDate: time.Now(), SendRequestDate: time.Now()})
		sysmailLog.MailitemID = uint(mailItemID)
		sysmailLog.LogDate = time.Now()
		sysmailLog.LastModDate = time.Now()
		sysmailLog.Description = err.Error()
		db.Model(&sysmailLog).Create(&sysmailLog)
	} else {
		db.Model(&sysmailMailitems).Where(models.SysmailMailItems{MailItemID: mailItemID}).Updates(models.SysmailMailItems{SentStatus: uint(sent), SentDate: time.Now(), LastModDate: time.Now(), SendRequestDate: time.Now()})
	}

	logging.Information("Updated sysmail_mailitems!")

}

//GetSysmailItemsCount returns a list of sysmailMailitems
func (t T) GetSysmailItemsCount(db *gorm.DB) int64 {

	logging.Information("Getting sysmail_mailitems count...")

	var sysmailMailitems models.SysmailMailItems
	var count int64

	result := db.Model(&sysmailMailitems).Where(map[string]interface{}{"sent_status": unsent}).Count(&count)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logging.Critical(result.Error)
		os.Exit(1)
	}

	logging.Information("Count for sysmail_mailitems: " + strconv.FormatInt(count, 10))

	return count
}
