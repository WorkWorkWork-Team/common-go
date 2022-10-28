package databasemysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

func NewDbConnection(config Config) (*sql.DB, error) {
	logger := newLogrusEntry(config)

	db, err := sql.Open(
		"mysql",
		config.Username+":"+config.Password+"@tcp("+config.Hostname+")/"+config.DatabaseName,
	)
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
