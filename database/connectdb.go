package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func DatabaseConnection() (*gorm.DB, error) {
	host := "ec2-44-207-253-50.compute-1.amazonaws.com"
	port := "5432"
	dbName := "deenk3a2vi1vi3"
	dbUser := "vplfguuhnqncri"
	password := "011ae171c243e0a6139df827e8393aef8954007ef25e552ca80c560e9ed3dc76"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	} else {
		fmt.Println("Database connection successful...")
	}
	return DB, err
}
