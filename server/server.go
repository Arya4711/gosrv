package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gosrv/util"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := gosrv.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	for pattern, handler := range gosrv.Router(db) {
		http.HandleFunc(pattern, handler)
	}

	log.Printf("listening on port %v\n", os.Getenv("PORT"))
	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
