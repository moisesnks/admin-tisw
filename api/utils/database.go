package utils

import (
	"admin/api/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// OpenDB abre la conexión con la base de datos y la devuelve
func OpenDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión a la base de datos: %v", err)
	}

	// Comprobar que la conexión con la base de datos es correcta
	if pingErr := db.Ping(); pingErr != nil {
		db.Close()
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %v", pingErr)
	}

	return db, nil
}
