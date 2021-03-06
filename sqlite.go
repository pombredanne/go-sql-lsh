package sqllsh

import "database/sql"

// NewSqliteLsh creates a new Sqlite3-backed LSH index.
// The caller is responsible for closing the database connection
// object.
func NewSqliteLsh(k, l int, tableName string, db *sql.DB) (*SqlLsh, error) {
	varFmt := func(i int) string {
		return "?"
	}
	createIndexFmt := "CREATE INDEX ht_%d ON %s ("
	lsh, err := newSqlLsh(k, l, tableName, db, varFmt, createIndexFmt)
	return lsh, err
}
