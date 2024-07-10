package gosrv

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectToDB() (db *sql.DB, e error) {
	connStr := fmt.Sprintf("postgresql://%v:%v@%v/test",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		fmt.Sprintf("localhost:%v",
			os.Getenv("PORT")))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		e = err
		return
	}

	return
}
