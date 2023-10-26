package storage

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

const DB_USERNAME = "myapiuser"
const DB_PASSWORD = "root"
const DB_NAME = "clothes_factory"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var DB *sql.DB

func InitDb() {
	cfg := mysql.Config{
		User:                 DB_USERNAME,
		Passwd:               DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 DB_HOST + ":" + DB_PORT,
		DBName:               DB_NAME,
		AllowNativePasswords: true,
		Collation:            "utf8mb4_general_ci",
	}

	// Get a driver-specific connector.
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Get a database handle.
	DB = sql.OpenDB(connector)
}
