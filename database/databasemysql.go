package databasemysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	IP           string
	Username     string
	Password     string
	DatabaseName string
}

func CreateConnection(config Config) *sql.DB {

	err := godotenv.Load(".env")

	db, err := sql.Open(
		"mysql",
		config.Username+":"+config.Password+"@tcp("+config.IP+")/"+config.DatabaseName,
	)
	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success!")

	return db
}
