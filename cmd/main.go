package main

import (
	"log"

	"github.com/conesForest/pingme_backend/internal/server"
	"github.com/conesForest/pingme_backend/pkg/db/postgres"
)

func main() {
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatalf("Can't connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Database created")

	s := server.New(db)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
