package database

import (
    "database/sql"
    "log"

    _ "modernc.org/sqlite" // Pure Go SQLite driver
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "./todo.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create table if not exists
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL,
        done BOOLEAN DEFAULT 0
    );`
    _, err = DB.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }
}
