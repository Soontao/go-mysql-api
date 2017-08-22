package inter

import (
	. "github.com/Soontao/go-mysql-api/t"
)

type ISQLBuilder interface {
	DeleteByTableAndId(tableName string, id interface{}) (sql string, err error)
	DeleteByTable(tableName string, mWhere map[string]interface{}) (sql string, err error)
	InsertByTable(tableName string, record map[string]interface{}) (sql string, err error)
	UpdateByTableAndId(tableName string, id interface{}, record map[string]interface{}) (sql string, err error)
	GetByTableAndID(opt QueryOption) (sql string, err error)
	GetByTable(opt QueryOption) (sql string, err error)
}
