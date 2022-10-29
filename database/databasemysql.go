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

func NewDbConnection(config Config) *sql.DB {

	db, err := sql.Open(
		"mysql",
		config.Username+":"+config.Password+"@tcp("+config.Hostname+")/"+config.DatabaseName,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Database": "MYSQL",
			"Hostname": config.Hostname,
		}).Error(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Database": "MYSQL",
			"Hostname": config.Hostname,
		}).Error(err)
	} else {
		logrus.WithFields(logrus.Fields{
			"Database": "MYSQL",
			"Hostname": config.Hostname,
		}).Info("Connected to Database")
	}

	return db
}
