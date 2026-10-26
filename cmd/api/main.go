package main

import (
	"fmt"
	"log"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	defer db.Close()

	fmt.Println("Successfully connected to LetsEat PostgreSQL!")
}
