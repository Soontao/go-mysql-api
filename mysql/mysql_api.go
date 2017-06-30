package mysql

import (
	"database/sql"
	"log"
	"time"
	// registe mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// MysqlAPI type
type MysqlAPI struct {
	connection *sql.DB
}

// GetConnectionPool which Pool is Singleton Connection Pool
func (api *MysqlAPI) GetConnectionPool(dbURI string) *sql.DB {
	if api.connection == nil {
		pool, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err.Error())
		}
		// 3 minutes unused connections will be closed
		pool.SetConnMaxLifetime(180 * time.Second)
		pool.SetMaxIdleConns(3)
		pool.SetMaxOpenConns(10)
		api.connection = pool
	}
	return api.connection
}

// Stop MysqlAPI, clean connections
func (api *MysqlAPI) Stop() *MysqlAPI {
	if api.connection != nil {
		api.connection.Close()
	}
	return api
}

// CurrentDatabaseName return current database
func (api *MysqlAPI) CurrentDatabaseName() string {
	rows, err := api.connection.Query("select database()")
	if err != nil {
		log.Fatal(err)
	}
	var res string
	for rows.Next() {
		if err := rows.Scan(&res); err != nil {
			log.Fatal(err)
		}
	}

	return res
}

func (api *MysqlAPI) retriveDatabaseMetadata() *MysqlAPI {

	return api
}

func (api *MysqlAPI) retriveTableMetadata() *MysqlAPI {

	return api
}

// NewMysqlAPI create new MysqlAPI instance
func NewMysqlAPI(dbURI string) *MysqlAPI {
	newAPI := &MysqlAPI{}
	newAPI.GetConnectionPool(dbURI)
	newAPI.retriveDatabaseMetadata()
	return newAPI
}
