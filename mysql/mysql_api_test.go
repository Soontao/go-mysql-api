package mysql

import (
	"log"
	"testing"
)

var coonectionStr = "root:stuecstothetimetolife@tcp(stu.ecs.fornever.org:3306)/monitor"

func TestCreateMysqlAPIInstance(t *testing.T) {
	api := NewMysqlAPI(coonectionStr)
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
	api := NewMysqlAPI(coonectionStr)
	res := api.CurrentDatabaseName()
	println(res)
	api.Stop()
}
