package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("server started at localhost:%v \n", os.Getenv("portApp"))
	http.ListenAndServe(os.Getenv("portApp"), nil)
}
