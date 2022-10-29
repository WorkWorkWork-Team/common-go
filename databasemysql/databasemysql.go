package databasemysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Hostname     string
	Username     string
	Password     string
	DatabaseName string
}

func newLogrusEntry(config Config) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Database": "MYSQL",
		"Hostname": config.Hostname,
	})
}

func NewDbConnection(config Config) (*sqlx.DB, error) {
	logger := newLogrusEntry(config)
	// Open MySQL connection.
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.Username, config.Password, config.Hostname, config.DatabaseName))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Info("Connected to Database")
	return db, nil
}
