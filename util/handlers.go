package gosrv

import (
	"database/sql"
	"net/http"
)

func Router(db *sql.DB) map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"GET /": func(w http.ResponseWriter, req *http.Request) { http.ServeFile(w, req, "util/test.html") },
	}
}
