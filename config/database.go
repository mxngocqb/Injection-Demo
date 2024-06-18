package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	// host     = "postgresql"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "injection"
)

func ConnectDatabase() *gorm.DB {
	fmt.Println("Connecting to the database...")

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		fmt.Println("Database connected")
	}

	return db
}
