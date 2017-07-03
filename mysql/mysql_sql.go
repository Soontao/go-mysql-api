package mysql

import (
	"gopkg.in/doug-martin/goqu.v4"
)

// SQL return sqls by sql builder
type SQL struct {
	sqlBuilder *goqu.Database
	dbMeta     *DataBaseMetadata
}

// GetByTableAndID sql
func (this *SQL) GetByTableAndID(tableName string, id interface{}) (sql string, err error) {
	table := this.dbMeta.GetTableMeta(tableName)
	priKey := table.GetPrimaryColumn().ColumnName
	sql, _, err = this.sqlBuilder.From(tableName).Where(goqu.Ex{priKey: id}).ToSql()
	return sql, err
}
