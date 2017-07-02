package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	// registe mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// MysqlAPI
type MysqlAPI struct {
	connection       *sql.DB
	databaseMetadata *DataBaseMetadata
}

// GetDatabaseMetadata return database meta
func (api *MysqlAPI) GetDatabaseMetadata() *DataBaseMetadata {
	return api.databaseMetadata
}

// GetConnectionPool which Pool is Singleton Connection Pool
func (api *MysqlAPI) GetConnectionPool(dbURI string) *sql.DB {
	if api.connection == nil {
		pool, err := sql.Open("mysql", dbURI)
		if err != nil {
			log.Fatal(err.Error())
		}
		// 3 minutes unused connections will be closed
		pool.SetConnMaxLifetime(3 * time.Minute)
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
	processIfError(err)
	var res string
	for rows.Next() {
		if err := rows.Scan(&res); err != nil {
			log.Fatal(err)
		}
	}
	return res
}

func (api *MysqlAPI) retriveDatabaseMetadata(databaseName string) *DataBaseMetadata {
	var tableMetas []*TableMetadata
	rs := &DataBaseMetadata{DatabaseName: databaseName}
	rows, err := api.connection.Query("show tables")
	processIfError(err)
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		processIfError(err)
		tableMetas = append(tableMetas, api.retriveTableMetadata(tableName))
	}
	rs.Tables = tableMetas
	return rs
}

func processIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (api *MysqlAPI) retriveTableMetadata(tableName string) *TableMetadata {
	rs := &TableMetadata{TableName: tableName}
	var columnMetas []*ColumnMetadata
	rows, err := api.connection.Query(fmt.Sprintf("desc %s", tableName))
	processIfError(err)
	for rows.Next() {
		var columnName, columnType, nullAble, key, defaultValue, extra sql.NullString
		err := rows.Scan(&columnName, &columnType, &nullAble, &key, &defaultValue, &extra)
		processIfError(err)
		columnMeta := &ColumnMetadata{columnName.String, columnType.String, nullAble.String, key.String, defaultValue.String, extra.String}
		columnMetas = append(columnMetas, columnMeta)
	}
	rs.Columns = columnMetas
	return rs
}

// NewMysqlAPI create new MysqlAPI instance
func NewMysqlAPI(dbURI string) *MysqlAPI {
	newAPI := &MysqlAPI{}
	newAPI.GetConnectionPool(dbURI)
	newAPI.databaseMetadata = newAPI.retriveDatabaseMetadata(newAPI.CurrentDatabaseName())
	return newAPI
}

// Query by sql
func (api *MysqlAPI) Query(sql string) ([]map[string]interface{}, error) {
	var rs []map[string]interface{}
	rows, _ := api.connection.Query(sql)
	cols, _ := rows.Columns()

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		rs = append(rs, m)
	}
	return rs, nil
}
