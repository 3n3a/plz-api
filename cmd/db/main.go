package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	ConnectionString string
}

func New(dbConfig Config) DB {
	db := DB{}
	return db.init(dbConfig.ConnectionString)
}

type DB struct {
	Conn *gorm.DB
}

func (db *DB) init(connString string) DB {
	dbConn, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to databse. \n", err)
	}

	log.Println("Connected to database.")

	return DB{
		Conn: dbConn,
	}
}