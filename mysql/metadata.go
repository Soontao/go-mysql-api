package mysql

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

// GetConnectionPool which Pool is Singleton Connection Pool
func GetConnectionPool(dbURI string) *sql.DB {
	if db == nil {
		pool, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err.Error())
		}
		// 3 minutes unused connections will be closed
		pool.SetConnMaxLifetime(180 * time.Second)
		pool.SetMaxIdleConns(3)
		pool.SetMaxOpenConns(10)
		db = pool
	}
	return db
}
