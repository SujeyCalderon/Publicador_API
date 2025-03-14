package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitDB inicializa la conexión a la base de datos.
func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}
