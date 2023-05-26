package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT IS NOT FOUND")
	}

	fmt.Println("Port:", portString)
}
