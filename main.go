package main

import (
	"log"

	"github.com/Zmahl/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPSServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
