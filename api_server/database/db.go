package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres",
		os.Getenv("DATABASE_USER")+":"+
			os.Getenv("DATABASE_PASSWORD")+"@tcp("+
			os.Getenv("DB_HOST")+":3306)/"+
			os.Getenv("DATABASE_NAME")+"?tls=false")
	if err != nil {
		panic(err)
	}

	return db
}
