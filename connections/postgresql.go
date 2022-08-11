package connections

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	PostgresqlConnection *sql.DB
}

var postgresql_url = "postgresql://admin:test@localhost:5432/test_db?sslmode=disable"

func (db *PostgreSQL) CreatePostgreSQLConnection() {
	connection, err := sql.Open("postgres", postgresql_url)

	if err != nil {
		log.Fatal(err)
	}

	db.PostgresqlConnection = connection
}

type SefaTable struct {
	ID   int
	Name string
}

func (db *PostgreSQL) PostgreSQLConnectionTest() {
	var mySefaTable SefaTable

	response, err := db.PostgresqlConnection.Query("SELECT * from sefa")

	if err != nil {
		log.Fatalln(err)
	}

	for response.Next() {
		err := response.Scan(&mySefaTable.ID, &mySefaTable.Name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(mySefaTable.ID, mySefaTable.Name)
	}
}
