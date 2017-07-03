package mysql

import (
	"net/url"
	"strconv"

	"gopkg.in/doug-martin/goqu.v4"
)

// SQL return sqls by sql builder
type SQL struct {
	sqlBuilder *goqu.Database
	dbMeta     *DataBaseMetadata
}

// GetByTableAndID for specific record in table
func (this *SQL) GetByTableAndID(tableName string, id interface{}, queryParam url.Values) (sql string, err error) {
	priKey := this.dbMeta.GetTableMeta(tableName).GetPrimaryColumn().ColumnName
	builder := this.sqlBuilder.From(tableName).Where(goqu.Ex{priKey: id})
	if queryParam["limit"] != nil {
		iLimit, err := strconv.ParseUint(queryParam["limit"][0], 10, 2)
		if err != nil {
			return "", err
		}
		builder.Limit(uint(iLimit))
	}
	sql, _, err = builder.ToSql()
	return sql, err
}

// GetByTable for most records
func (this *SQL) GetByTable(tableName string, queryParam url.Values) (sql string, err error) {
	builder := this.sqlBuilder.From(tableName)
	builder, err = processSQLBuilderWithQueryParam(builder, queryParam)
	sql, _, err = builder.ToSql()
	return
}

func processSQLBuilderWithQueryParam(builder *goqu.Dataset, queryParam url.Values) (rs *goqu.Dataset, err error) {
	rs = builder
	if queryParam["_limit"] != nil {
		iLimit, err := strconv.ParseUint(queryParam["_limit"][0], 10, 2)
		if err != nil {
			return rs, err
		}
		rs = rs.Limit(uint(iLimit))
	}
	if queryParam["_field"] != nil {
		fields := make([]interface{}, len(queryParam["_field"]))
		for idx, f := range queryParam["_field"] {
			fields[idx] = f
		}
		rs = rs.Select(fields...)
	}
	return rs, err
}
