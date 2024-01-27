package main

import (
	"log"

	"github.com/gabeno/logspace/internal/server"
)

func main() {
	server := server.NewHTTPServer(":8080")
	log.Fatal(server.ListenAndServe())
}
