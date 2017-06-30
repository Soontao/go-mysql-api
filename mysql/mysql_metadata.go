package mysql

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
