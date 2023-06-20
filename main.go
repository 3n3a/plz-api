package main

import (
	"log"
	"os"

	"github.com/3n3a/plz-api/cmd/db"
	"github.com/3n3a/plz-api/cmd/routes"
)

// @title Postal Codes API
// @version 1.0
// @description Fast and modern REST-API for Swiss Postal Codes
func main() {
	dbInstance := db.New(db.Config{
		ConnectionString: os.Getenv("DB_CONN"),
	})

    app := routes.New(dbInstance)

    log.Fatal(app.Listen(":3000"))
}