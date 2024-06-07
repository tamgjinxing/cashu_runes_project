package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager(dataSourceName string) (*DBManager, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DBManager{DB: db}, nil
}

func (manager *DBManager) Close() {
	if manager.DB != nil {
		manager.DB.Close()
	}
}
