package connections

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var PostgresqlConnection *sql.DB

const postgresql_url = "postgresql://admin:test@localhost:5432/test_db?sslmode=disable"

func CreatePostgreSQLConnection() {
	connection, err := sql.Open("postgres", postgresql_url)

	if err != nil {
		log.Fatal(err)
	}

	err = connection.Ping()

	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}

	log.Println("PostgreSQL Connected !")

	PostgresqlConnection = connection
}
