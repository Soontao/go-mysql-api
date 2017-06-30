package mysql

// DataBaseMetadata metadata of a database
type DataBaseMetadata struct {
	databaseName string          // database name
	tables       []TableMetadata // collection of tables
}

// TableMetadata metadata of a table
type TableMetadata struct {
	tableName string           //table name
	columns   []ColumnMetadata //collections of column
}

// ColumnMetadata metadata of a column
type ColumnMetadata struct {
	columnName    string // column name or code ?
	columnType    string // column type
	columnTypeLen int    // column type length, if have
	// If Key is PRI, the column is a PRIMARY KEY or is one of the
	// columns in a multiple-column PRIMARY KEY.

	// If Key is UNI, the column is the first column of a unique-valued
	// index that cannot contain NULL values.

	// If Key is MUL, multiple occurrences of a given value are
	// permitted within the column. The column is the first column
	// of a nonunique index or a unique-valued index that can contain
	// NULL values.
	key          string // column key type
	defaultValue string // default value if have
	extra        string // extra info, for example, auto_increment
}
