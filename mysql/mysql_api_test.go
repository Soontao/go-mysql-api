package mysql

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

var connectionStr = os.Getenv("API_CONN_STR")

func TestCreateMysqlAPIInstance(t *testing.T) {
	api := NewMysqlAPI(connectionStr, true)
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
	api := NewMysqlAPI(connectionStr, true)
	res := api.CurrentDatabaseName()
	println(res)
	api.Stop()
}

func TestRetriveMetadata(t *testing.T) {
	api := NewMysqlAPI(connectionStr, true)
	res := api.retriveDatabaseMetadata("monitor")
	println(res.Tables[0].Columns[0].ColumnName)
	api.Stop()
}

func TestRowScan(t *testing.T) {
	api := NewMysqlAPI(connectionStr, true)
	defer api.Stop()
	rs, err := api.query("select * from monitor Limit ?", 2)
	if err != nil {
		t.Error(err)
	}
	for _, row := range rs {
		jsonStr, _ := json.Marshal(row) // use gojson avoid base64 encode of []byte
		fmt.Printf("%s\n", jsonStr)
	}
}
