package gosrv

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Router(db *sql.DB) map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"GET /users": func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			var res User
			var users []User
			rows, err := db.Query("SELECT * FROM users")
			defer rows.Close()
			if err != nil {
				log.Fatalln(err)
				fmt.Fprintln(w, "an error occurred")
			}

			for rows.Next() {
				rows.Scan(&res)
				users = append(users, res)
			}

			json.NewEncoder(w).Encode(users)
		},
	}
}
