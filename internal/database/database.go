package database

import (
	"database/sql"
	"log"

	"github.com/jwDevOps/atlas-backend/internal/models"
)

type DbManager struct {
	Db *sql.DB
}

func (m *DbManager) CreateSchema() {
	if _, err := m.Db.Exec(models.PUBLISHER_SCHEMA); err != nil {
		log.Fatal(err)
	}
}

func Init(dbName string) *DbManager {
	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatal(err)
	}

	manager := new(DbManager)
	manager.Db = db

	return manager
}
