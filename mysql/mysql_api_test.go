package mysql

import (
	"log"
	"testing"
)

var connectionStr = "monitor:yn0Mbx1mPcZWlvzb@tcp(stu.ecs.fornever.org:3306)/monitor"

func TestCreateMysqlAPIInstance(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	rows, err := api.connection.Query("select 1")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var col1 string
		if err := rows.Scan(&col1); err != nil {
			log.Fatal(err)
		}
		println(col1)
	}
	api.Stop()
}

func TestCurrentDatabaseName(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	res := api.CurrentDatabaseName()
	println(res)
	api.Stop()
}

func TestRetriveMetadata(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	res := api.retriveDatabaseMetadata("monitor")
	println(res.Tables[0].Columns[0].ColumnName)
	api.Stop()
}
