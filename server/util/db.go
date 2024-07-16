package gosrv

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type User struct {
	id       int
	username string
	password string
}

func ConnectToDB() (db *sql.DB, e error) {
	connStr := fmt.Sprintf("postgresql://%v:%v@%v/test?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		fmt.Sprintf("%v:%v", os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_PORT")),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e = err
		return
	}

	return
}
