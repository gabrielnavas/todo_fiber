package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func MakeDatabase(username, password, host, port, databaseName, sslmode string) *sql.DB {
	var err error

	connStr := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + databaseName + "?sslmode=" + sslmode
	// connStr := "postgres://postgres:password123@localhost:5435/task?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
