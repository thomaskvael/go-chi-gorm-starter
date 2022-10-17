package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thomaskvael/go-chi-gorm-starter/router"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	fmt.Println("Hello world!")
	http.ListenAndServe(":3000", router.RegisterRouter())
}
