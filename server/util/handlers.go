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
			rows, err := db.Query("SELECT * FROM users")
			defer rows.Close()
			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to query database")
			}

			var user User
			var users []User
			for rows.Next() {
				if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
					log.Fatalln(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w, "failed to scan")
				}
				users = append(users, user)
			}

			if err = json.NewEncoder(w).Encode(users); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		},
		"POST /users": func(w http.ResponseWriter, req *http.Request) {
			var user User
			if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to read body")
			}

			if _, err := db.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%v', '%v')", user.Username, user.Password)); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to insert user")
			}
		},
	}
}
