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

func (s *SQL) getPriKeyNameOf(tableName string) string {
	return s.dbMeta.GetTableMeta(tableName).GetPrimaryColumn().ColumnName
}

// GetByTableAndID for specific record in table
func (s *SQL) GetByTableAndID(tableName string, id interface{}, queryParam url.Values) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: id})
	processSQLBuilderWithQueryParam(builder, queryParam)
	sql, _, err = builder.ToSql()
	return sql, err
}

// UpdateByTableAndID for update specific record by id
func (s *SQL) UpdateByTableAndID(tableName string, record map[string]interface{}) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: record[priKeyName]})
	sql, _, err = builder.ToUpdateSql(record)
	return
}

// InsertByTable and record map
func (s *SQL) InsertByTable(tableName string, record map[string]interface{}) (sql string, err error) {
	sql = goqu.From(tableName).Insert(record).Sql
	return
}

// DeleteByTableWhere map
func (s *SQL) DeleteByTableWhere(tableName string, mWhere map[string]interface{}) (sql string, err error) {
	builder := goqu.From(tableName)
	for k, v := range mWhere {
		builder = builder.Where(goqu.Ex{k: v})
	}
	sql = builder.Delete().Sql
	return
}

// GetByTable for most records
func (s *SQL) GetByTable(tableName string, queryParam url.Values) (sql string, err error) {
	builder := s.sqlBuilder.From(tableName)
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
