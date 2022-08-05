package postgres

import (
	"database/sql"
	"os"

	"fmt"

	_ "github.com/lib/pq"
)

var (
	hostname      = os.Getenv("POSTGRES_HOST")
	host_port     = 5432
	username      = os.Getenv("POSTGRES_USER")
	password      = os.Getenv("POSTGRES_PASSWORD")
	database_name = os.Getenv("POSTGRES_DB")
)

func Connect() *sql.DB {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		fmt.Printf("Error conecting to the DB", err)
	}

	// db.Exec("CREATE TABLE IF NOT EXISTS registration (number_plate VARCHAR(20) NOT NULL, vehicle JSON NOT NULL, insurer JSON NOT NULL);")

	return db
}
