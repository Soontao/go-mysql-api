package mysql

import (
	"gopkg.in/doug-martin/goqu.v4"
)

// DataBaseMetadata metadata of a database
type DataBaseMetadata struct {
	DatabaseName string           // database name
	Tables       []*TableMetadata // collection of tables
}

// TableMetadata metadata of a table
type TableMetadata struct {
	TableName string            //table name
	Columns   []*ColumnMetadata //collections of column
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
	Key          string // column key type
	DefaultValue string // default value if have
	Extra        string // extra info, for example, auto_increment
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
func (d *DataBaseMetadata) GetSimpleMetadata() (rt map[string]map[string]string) {
	rt = make(map[string]map[string]string)
	for _, table := range d.Tables {
		t := make(map[string]string)
		for _, f := range table.Columns {
			t[f.ColumnName] = f.ColumnType
		}
		rt[table.TableName] = t
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
