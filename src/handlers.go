package gosrv

import "net/http"

var Router = map[string]func(http.ResponseWriter, *http.Request){
	"GET /": func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "src/test.html")
	},
}
