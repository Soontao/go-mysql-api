package mysql

import "testing"
import "net/url"

func TestSQL_GetByTableAndID(t *testing.T) {
	api := NewMysqlAPI(connectionStr)
	defer api.Stop()
	if sql, err := api.sql.GetByTableAndID("monitor", 1, url.Values{}); err != nil {
		t.Error(err)
	} else {
		println(sql)
	}
}
