package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Database struct {
	dbPath    string
	dbPointer *sql.DB
}

var (
	DB                 *Database = &Database{}
	getTableNamesQuery           = `
	SELECT name FROM sqlite_schema
	WHERE type = 'table'
	AND name NOT LIKE 'sqlite_%';`
)

func (d *Database) Init(dbPath string) {
	d.dbPath = dbPath
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	d.dbPointer = db
}

func (d *Database) GetDatabaseTables() []string {
	rows, err := d.dbPointer.Query(getTableNamesQuery, nil)
	if err != nil {
		log.Fatal(err)
	}
	tables := make([]string, 0)
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tables = append(tables, name)
	}

	log.Printf("%v", tables)
	return tables
}
