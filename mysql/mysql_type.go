package mysql

import (
	"gopkg.in/doug-martin/goqu.v4"
	"fmt"
)

// DataBaseMetadata metadata of a database
type DataBaseMetadata struct {
	DatabaseName string           // database name
	Tables       []*TableMetadata // collection of tables
}

// TableMetadata metadata of a table
type TableMetadata struct {
	TableName    string //table name
	TableType    string
	TableRows    int64
	CurrentIncre int64
	Comment      string
	Columns      []*ColumnMetadata //collections of column
}

// ColumnMetadata metadata of a column
type ColumnMetadata struct {
	ColumnName string // column name or code ?
	ColumnType string // column type
	NullAble   string // column null able
	// If Key is PRI, the column is a PRIMARY KEY or is one of the
	// columns in a multiple-column PRIMARY KEY.

	// If Key is UNI, the column is the first column of a unique-valued
	// index that cannot contain NULL values.

	// If Key is MUL, multiple occurrences of a given value are
	// permitted within the column. The column is the first column
	// of a nonunique index or a unique-valued index that can contain
	// NULL values.
	Key              string // column key type
	DefaultValue     string // default value if have
	Extra            string // extra info, for example, auto_increment
	OridinalSequence int64
	DataType         string
	Comment          string
}

// QueryConfig for Select method
type QueryOption struct {
	limit  int
	offset int
	fields []interface{}
	links  []interface{}
	wheres map[string]goqu.Op
}

// GetTableMeta
func (d *DataBaseMetadata) GetTableMeta(tableName string) *TableMetadata {
	for _, table := range d.Tables {
		if table.TableName == tableName {
			return table
		}
	}
	return nil
}

// GetSimpleMetadata
func (d *DataBaseMetadata) GetSimpleMetadata() (rt map[string]interface{}) {
	rt = make(map[string]interface{})
	for _, table := range d.Tables {
		cls := make([]interface{}, len(table.Columns))
		for idx, i_column := range table.Columns {
			cls[idx] = fmt.Sprintf("%s %s %s NullAble(%s) '%s'", i_column.ColumnName, i_column.ColumnType, i_column.DefaultValue, i_column.NullAble, i_column.Comment)
		}
		rt[fmt.Sprintf("[%s] (%d rows) %s", table.TableType, table.TableRows, table.TableName)] = cls
	}
	return
}

// GetPrimaryColumn
func (t *TableMetadata) GetPrimaryColumn() *ColumnMetadata {
	for _, col := range t.Columns {
		if col.Key == "PRI" {
			return col
		}
	}
	return nil
}

// HaveField
func (t *TableMetadata) HaveField(sFieldName string) bool {
	for _, col := range t.Columns {
		if sFieldName == col.ColumnName {
			return true
		}
	}
	return false
}

// HaveTable
func (d *DataBaseMetadata) HaveTable(sTableName string) bool {
	if t := d.GetTableMeta(sTableName); t != nil {
		return true
	}
	return false
}

// TableHaveField
func (d *DataBaseMetadata) TableHaveField(sTableName string, sFieldName string) bool {
	t := d.GetTableMeta(sTableName)
	if t == nil {
		return false
	}
	return t.HaveField(sFieldName)
}
