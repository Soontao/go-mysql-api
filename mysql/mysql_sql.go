package mysql

import (
	"fmt"

	"gopkg.in/doug-martin/goqu.v4"
	. "github.com/Soontao/go-mysql-api/types"
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
func (s *SQL) GetByTable(opt QueryOption) (sql string, err error) {
	builder := s.sqlBuilder.From(opt.Table)
	builder = s.configBuilder(builder, opt.Table, opt)
	sql, _, err = builder.ToSql()
	return
}

// GetByTableAndID for specific record in Table
func (s *SQL) GetByTableAndID(opt QueryOption) (sql string, err error) {
	priKeyName := s.getPriKeyNameOf(opt.Table)
	builder := s.sqlBuilder.From(opt.Table).Where(goqu.Ex{priKeyName: opt.Id})
	builder = s.configBuilder(builder, opt.Table, opt)
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
		err = fmt.Errorf("Table `%s` dont have primary key !/n", tableName)
		return
	} else {
		return s.DeleteByTable(tableName, map[string]interface{}{priKeyName: id})
	}
}

func (s *SQL) configBuilder(builder *goqu.Dataset, priT string, opt QueryOption) (rs *goqu.Dataset) {
	rs = builder
	if opt.Limit != 0 {
		rs = rs.Limit(uint(opt.Limit))
	}
	if opt.Offset != 0 {
		rs = rs.Offset(uint(opt.Offset))
	}
	if opt.Fields != nil {
		fs := make([]interface{}, len(opt.Fields))
		for idx, f := range opt.Fields {
			fs[idx] = f
		}
		rs = rs.Select(fs...)
	}
	for f, w := range opt.Wheres {
		// check field exist
		rs = rs.Where(goqu.Ex{f: goqu.Op{w.Operation: w.Value}})
	}
	for _, l := range opt.Links {
		refT := l
		refK := s.getPriKeyNameOf(refT)
		priK := s.getPriKeyNameOf(priT)
		if s.dbMeta.TableHaveField(priT, refK) {
			rs = rs.InnerJoin(goqu.I(refT), goqu.On(goqu.I(fmt.Sprintf("%s.%s", refT, refK)).Eq(goqu.I(fmt.Sprintf("%s.%s", priT, refK)))))
		}
		if s.dbMeta.TableHaveField(refT, priK) {
			rs = rs.InnerJoin(goqu.I(refT), goqu.On(goqu.I(fmt.Sprintf("%s.%s", refT, priK)).Eq(goqu.I(fmt.Sprintf("%s.%s", priT, priK)))))
		}
	}
	if opt.Search != "" {
		searchEx := goqu.ExOr{}
		for _, c := range s.dbMeta.GetTableMeta(opt.Table).Columns {
			searchEx[c.ColumnName] = goqu.Op{"like": fmt.Sprintf("%%%s%%", opt.Search)}
		}
		rs = rs.Where(searchEx)
	}
	return
}
