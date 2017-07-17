package mysql

import (
	"fmt"

	"gopkg.in/doug-martin/goqu.v4"

	_ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

// SQL return sqls by sql builder
type SQL struct {
	sqlBuilder *goqu.Database
	dbMeta     *DataBaseMetadata
}

func (s *SQL) getPriKeyNameOf(tableName string) string {
	return s.dbMeta.GetTableMeta(tableName).GetPrimaryColumn().ColumnName
}

// GetByTable with filter
func (s *SQL) GetByTable(tableName string, opt QueryOption) (sql string, err error) {
	builder := s.sqlBuilder.From(tableName)
	builder = s.configBuilder(builder, tableName, opt)
	sql, _, err = builder.ToSql()
	return
}

// GetByTableAndID for specific record in table
func (s *SQL) GetByTableAndID(tableName string, id interface{}, opt QueryOption) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: id})
	builder = s.configBuilder(builder, tableName, opt)
	sql, _, err = builder.ToSql()
	return sql, err
}

// UpdateByTable for update specific record by id
func (s *SQL) UpdateByTableAndId(tableName string, id interface{}, record map[string]interface{}) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	builder := s.sqlBuilder.From(tableName).Where(goqu.Ex{priKeyName: id})
	sql, _, err = builder.ToUpdateSql(record)
	return
}

// InsertByTable and record map
func (s *SQL) InsertByTable(tableName string, record map[string]interface{}) (sql string, err error) {
	sql, _, err = s.sqlBuilder.From(tableName).Where().ToInsertSql(record)
	return
}

// DeleteByTable by where
func (s *SQL) DeleteByTable(tableName string, mWhere map[string]interface{}) (sql string, err error) {
	builder := s.sqlBuilder.From(tableName)
	for k, v := range mWhere {
		builder = builder.Where(goqu.Ex{k: v})
	}
	sql = builder.Delete().Sql
	return
}

// DeleteByTableAndId
func (s *SQL) DeleteByTableAndId(tableName string, id interface{}) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(tableName)
	if priKeyName == "" {
		err = fmt.Errorf("table `%s` dont have primary key !/n", tableName)
		return
	} else {
		return s.DeleteByTable(tableName, map[string]interface{}{priKeyName: id})
	}
}

func (s *SQL) configBuilder(builder *goqu.Dataset, tName string, opt QueryOption) (rs *goqu.Dataset) {
	rs = builder
	if opt.limit != 0 {
		rs = rs.Limit(uint(opt.limit))
	}
	if opt.offset != 0 {
		rs = rs.Offset(uint(opt.offset))
	}
	if opt.fields != nil {
		rs = rs.Select(opt.fields...)
	}
	for _, w := range opt.wheres {
		rs = rs.Where(goqu.Ex{w.Field.(string): w.Operator})
	}
	for _, l := range opt.links {
		refT := l.(string)
		refK := s.getPriKeyNameOf(refT)
		priK := s.getPriKeyNameOf(tName)
		if s.dbMeta.TableHaveField(tName, refK) {
			rs = rs.InnerJoin(goqu.I(refT), goqu.On(goqu.I(fmt.Sprintf("%s.%s", refT, refK)).Eq(goqu.I(fmt.Sprintf("%s.%s", tName, refK)))))
		}
		if s.dbMeta.TableHaveField(refT, priK) {
			rs = rs.InnerJoin(goqu.I(refT), goqu.On(goqu.I(fmt.Sprintf("%s.%s", refT, priK)).Eq(goqu.I(fmt.Sprintf("%s.%s", tName, priK)))))
		}
	}
	return
}
