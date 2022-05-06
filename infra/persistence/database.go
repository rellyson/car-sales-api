package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rellyson/car-sales-api/application/utils"
)

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	DbName   string
	SSLMode  string
}

var (
	db     *sql.DB
	logger = utils.NewLogger()
)

func CreateDBConnection(c DatabaseConfig) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DbName)

	dbConn, err := sql.Open(c.Driver, connStr)

	if err != nil {
		logger.Error(fmt.Sprintf("there was an error openning database: %s", err.Error()))
	}

	db = dbConn

	//pings database to check connectivity
	if err = db.Ping(); err != nil {
		logger.Error(fmt.Sprintf("there was an error when trying to reach db: %s", err.Error()))
	}
}

func GetDBConnection() *sql.DB {
	return db
}
