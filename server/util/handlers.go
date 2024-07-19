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
		"GET /users/{username}": func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			username := req.PathValue("username")
			if username == "" {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to retrieve username")
			}

			query, err := db.Query(fmt.Sprintf("SELECT * FROM users WHERE username=%v", username))
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to find user")
			}

			var user User
			if err := query.Scan(&user.Id, &user.Username, &user.Password); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to scan")
			}

			if err = json.NewEncoder(w).Encode(user); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to send JSON")
			}
		},
		"GET /users": func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			rows, err := db.Query("SELECT * FROM users")
			defer rows.Close()
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to query database")
			}

			var user User
			var users []User
			for rows.Next() {
				if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w, "failed to scan")
				}
				users = append(users, user)
			}

			if err = json.NewEncoder(w).Encode(users); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to send JSON")
			}
		},
		"POST /users": func(w http.ResponseWriter, req *http.Request) {
			var user User
			if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to read body")
			}

			if _, err := db.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%v', '%v')", user.Username, user.Password)); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to insert user")
			}
		},
		"DELETE /users": func(w http.ResponseWriter, req *http.Request) {
			var temp User
			if err := json.NewDecoder(req.Body).Decode(&temp); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to read body")
			}

			if _, err := db.Exec(fmt.Sprintf("DELETE FROM users WHERE username='%v'", temp.Username)); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "failed to delete user")
			}

			w.WriteHeader(http.StatusNoContent)
		},
	}
}
