package main

import (
	"database/sql"
	"fmt"
	"github.com/bastilavarias/frontzero/core"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	manifest, _ := core.LoadManifest("frontzero.yaml")
	dbCredentials := core.DatabaseCredentials{
		Path:   "./frontzero.db",
		Driver: "sqlite3",
	}
	db, err := sql.Open(dbCredentials.Driver, dbCredentials.Path)
	if err != nil {
		fmt.Println(err)
	}
	var columns []core.DatabaseTableColumn
	for _, entity := range manifest.Entities {
		for _, field := range entity.Fields {
			columns = append(columns, core.DatabaseTableColumn{
				Name: entity.Name,
				Type: field.Type,
			})
		}
		table := core.DatabaseTable{
			Name:    entity.Name,
			Columns: columns,
		}
		err := core.GenerateDatabaseSchema(db, table)
		if err != nil {
			return
		}
	}
	schemas, _ := core.GetDatabaseSchemas(db)
	for _, schema := range schemas {
		fmt.Println(schema)
	}
}
