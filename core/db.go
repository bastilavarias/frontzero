package core

import (
	"database/sql"
	"fmt"
	"strings"
)

func GenerateDatabaseSchema(db *sql.DB, table DatabaseTable) error {
	var cols []string
	for _, col := range table.Columns {
		cols = append(cols, fmt.Sprintf("`%s` %s", col.Name, col.Type))
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s)", table.Name, strings.Join(cols, ","))
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}

	// Check if the table exists -> Check the columns (type, properties) already exists.

	return err
}

func GetDatabaseSchemas(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}
