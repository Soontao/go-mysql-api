package mysql

import (
	"fmt"
)

// DataBaseMetadata metadata of a database
type DataBaseMetadata struct {
	DatabaseName string           `json:"database_name,omitempty"` // database name
	Tables       []*TableMetadata `json:"tables,omitempty"`        // collection of tables
}

// TableMetadata metadata of a Table
type TableMetadata struct {
	TableName    string            `json:"table_name,omitempty"` //Table name
	TableType    string            `json:"table_type,omitempty"`
	TableRows    int64             `json:"table_rows,omitempty"`
	CurrentIncre int64             `json:"current_increment,omitempty"`
	Comment      string            `json:"comment,omitempty"`
	Columns      []*ColumnMetadata `json:"columns,omitempty"` //collections of column
}

// ColumnMetadata metadata of a column
type ColumnMetadata struct {
	ColumnName string `json:"column_name,omitempty"` // column name or code ?
	ColumnType string `json:"column_type,omitempty"` // column type
	NullAble   string `json:"nullable,omitempty"`    // column null able
	// If Key is PRI, the column is a PRIMARY KEY or is one of the
	// columns in a multiple-column PRIMARY KEY.

	// If Key is UNI, the column is the first column of a unique-valued
	// index that cannot contain NULL values.

	// If Key is MUL, multiple occurrences of a given value are
	// permitted within the column. The column is the first column
	// of a nonunique index or a unique-valued index that can contain
	// NULL values.
	Key              string `json:"key,omitempty"`           // column key type
	DefaultValue     string `json:"default_value,omitempty"` // default value if have
	Extra            string `json:"extra,omitempty"`         // extra info, for example, auto_increment
	OridinalSequence int64  `json:"oridinal_sequence,omitempty"`
	DataType         string `json:"data_type,omitempty"`
	Comment          string `json:"comment,omitempty"`
}

// QueryConfig for Select method
type QueryOption struct {
	Table  string                    // table name
	Id     string                    // select with primary key value
	Limit  int                       // record limit
	Offset int                       // start offset
	Fields []string                  // select fields
	Links  []string                  // auto join table
	Wheres map[string]WhereOperation // field -> { operation, value }
	Search string                    // fuzzy query word
}

type WhereOperation struct {
	Operation string
	Value     interface{}
}

func (c *ColumnMetadata) GetDefaultValue() (v interface{}) {
	if c.DefaultValue != "" {
		v = c.DefaultValue
	}
	return
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
