package mysql

import (
	"testing"

	goqu "gopkg.in/doug-martin/goqu.v4"
	// mysql dialect
	_ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

func TestSQL_GetByTableAndID(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	f := []interface{}{"mid"}
	defer api.Stop()
	if sql, err := api.sql.GetByTableAndID("monitor", 1, QueryOption{1, 2, f}); err != nil {
		t.Error(err)
	} else {
		println(sql)
	}
}

func TestInsertSQLFromMap(t *testing.T) {
	m := map[string]interface{}{"name": "monitor", "seq": 1}
	s := goqu.From("Table").Insert(m).Sql
	println(s)
}

func TestDeleteSQLFromMap(t *testing.T) {
	m := map[string]interface{}{"name": "monitor", "seq": 1}
	builder := goqu.From("DTable")
	for k, v := range m {
		builder = builder.Where(goqu.Ex{k: v})
	}
	s := builder.Delete().Sql
	println(s)
}

func TestUpdateSQLFromMap(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	defer api.Stop()
	if sql, err := api.sql.UpdateByTable("monitor", map[string]interface{}{"mid": 1, "target": "change it"}); err != nil {
		t.Error(err)
	} else {
		println(sql)
	}
}
