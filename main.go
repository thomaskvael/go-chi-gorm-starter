package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	fmt.Println("Hello world!")
}
