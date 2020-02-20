package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "service ", log.LstdFlags|log.Lshortfile)

	h := home.NewHandlers(logger)

	router := mux.NewRouter()
	router.HandleFunc("/", h.Page)

	logger.Println("Server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
